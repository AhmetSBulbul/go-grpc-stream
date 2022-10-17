package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/AhmetSBulbul/go-grpc-stream/protos"
	"google.golang.org/grpc"
)

type commsServiceServer struct {
	protos.UnimplementedCommsServiceServer
	channel map[string][]chan *protos.Message
}

func (s *commsServiceServer) JoinChannel(ch *protos.Channel, msgStream protos.CommsService_JoinChannelServer) error {
	msgChannel := make(chan *protos.Message)
	s.channel[ch.ChannelId] = append(s.channel[ch.ChannelId], msgChannel)
	fmt.Printf("Channel %s joined\n", ch.ChannelId)

	// doing this never closes the stream
	for {
		select {
		case <-msgStream.Context().Done():
			return nil
		case msg := <-msgChannel:
			fmt.Printf("GO ROUTINE (got message): %v \n", msg)
			msgStream.Send(msg)
		}
	}
}

func (s *commsServiceServer) SendMessage(msgStream protos.CommsService_SendMessageServer) error {
	msg, err := msgStream.Recv()

	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}

	ack := protos.MessageAck{Status: "OK"}
	msgStream.SendAndClose(&ack)

	go func() {
		streams := s.channel[msg.Channel.ChannelId]
		for _, msgChan := range streams {
			msgChan <- msg
		}
	}()

	return nil
}

func newServer() *commsServiceServer {
	return &commsServiceServer{
		channel: make(map[string][]chan *protos.Message),
	}
}

func main() {
	fmt.Println("--- SERVER APP ---")
	lis, err := net.Listen("tcp", "0.0.0.0:5400")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	protos.RegisterCommsServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
