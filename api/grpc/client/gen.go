package client

//go:generate protoc --go_out=./specification  --go-grpc_out=./specification  --proto_path=./../ EventService.proto
