gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/commspb.proto


install:
	go install .
	which go-grpc-stream

run:
	go run .