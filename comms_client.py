import threading
import time
import grpc
import commspb_pb2
import commspb_pb2_grpc


def get_messages():
    while True:
        # message = input("Please enter a message (or nothing to stop chatting): ")

        # if message == "":
            # break

        room = commspb_pb2.Channel(channel_id = 'test')
        status = commspb_pb2.Message(channel = room, message = "Hey")

        yield status
        time.sleep(1)


def __listen_for_messages(stub, room):
    for message in stub.JoinChannel(room):
        print(room + ": " + message)

def run():
    with grpc.insecure_channel('0.0.0.0:5400') as channel:
        stub = commspb_pb2_grpc.CommsServiceStub(channel)
        room = commspb_pb2.Channel(channel_id = 'test')
        threading.Thread(target=__listen_for_messages, args=(stub, room)).start()
        # for message in response:
        #     print(message)
        # print(status)
        # stub.SendMessage(get_messages())
        

        
        # response.listen(print)

if __name__ == "__main__":
    run()