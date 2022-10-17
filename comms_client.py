import time
import grpc
import commspb_pb2
import commspb_pb2_grpc


def get_messages():
    while True:
        message = input("Please enter a message (or nothing to stop chatting): ")

        if message == "":
            break

        room = commspb_pb2.Channel(channel_id = 'test')
        status = commspb_pb2.Message(channel = room, message = message)

        yield status
        time.sleep(1)




def run():
    with grpc.insecure_channel('0.0.0.0:5400') as channel:
        stub = commspb_pb2_grpc.CommsServiceStub(channel)
        room = commspb_pb2.Channel(channel_id = 'test')
        response = stub.JoinChannel(room)
        for message in response:
            print(message)
        status = stub.SendMessage(get_messages())
        print(status)
        

        
        # response.listen(print)

if __name__ == "__main__":
    run()