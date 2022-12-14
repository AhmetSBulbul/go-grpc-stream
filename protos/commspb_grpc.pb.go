// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: protos/commspb.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommsServiceClient is the client API for CommsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommsServiceClient interface {
	JoinChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (CommsService_JoinChannelClient, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (CommsService_SendMessageClient, error)
}

type commsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommsServiceClient(cc grpc.ClientConnInterface) CommsServiceClient {
	return &commsServiceClient{cc}
}

func (c *commsServiceClient) JoinChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (CommsService_JoinChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &CommsService_ServiceDesc.Streams[0], "/commspb.CommsService/JoinChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &commsServiceJoinChannelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommsService_JoinChannelClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type commsServiceJoinChannelClient struct {
	grpc.ClientStream
}

func (x *commsServiceJoinChannelClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *commsServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (CommsService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CommsService_ServiceDesc.Streams[1], "/commspb.CommsService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &commsServiceSendMessageClient{stream}
	return x, nil
}

type CommsService_SendMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*MessageAck, error)
	grpc.ClientStream
}

type commsServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *commsServiceSendMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commsServiceSendMessageClient) CloseAndRecv() (*MessageAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommsServiceServer is the server API for CommsService service.
// All implementations must embed UnimplementedCommsServiceServer
// for forward compatibility
type CommsServiceServer interface {
	JoinChannel(*Channel, CommsService_JoinChannelServer) error
	SendMessage(CommsService_SendMessageServer) error
	mustEmbedUnimplementedCommsServiceServer()
}

// UnimplementedCommsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommsServiceServer struct {
}

func (UnimplementedCommsServiceServer) JoinChannel(*Channel, CommsService_JoinChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinChannel not implemented")
}
func (UnimplementedCommsServiceServer) SendMessage(CommsService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedCommsServiceServer) mustEmbedUnimplementedCommsServiceServer() {}

// UnsafeCommsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommsServiceServer will
// result in compilation errors.
type UnsafeCommsServiceServer interface {
	mustEmbedUnimplementedCommsServiceServer()
}

func RegisterCommsServiceServer(s grpc.ServiceRegistrar, srv CommsServiceServer) {
	s.RegisterService(&CommsService_ServiceDesc, srv)
}

func _CommsService_JoinChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommsServiceServer).JoinChannel(m, &commsServiceJoinChannelServer{stream})
}

type CommsService_JoinChannelServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type commsServiceJoinChannelServer struct {
	grpc.ServerStream
}

func (x *commsServiceJoinChannelServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _CommsService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommsServiceServer).SendMessage(&commsServiceSendMessageServer{stream})
}

type CommsService_SendMessageServer interface {
	SendAndClose(*MessageAck) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type commsServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *commsServiceSendMessageServer) SendAndClose(m *MessageAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commsServiceSendMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommsService_ServiceDesc is the grpc.ServiceDesc for CommsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commspb.CommsService",
	HandlerType: (*CommsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinChannel",
			Handler:       _CommsService_JoinChannel_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _CommsService_SendMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protos/commspb.proto",
}
