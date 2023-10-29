package specification

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

//go:generate protoc --go_out=./internal/server/grpc/specification  --go-grpc_out=./internal/server/grpc/specification --go-grpc_out=require_unimplemented_servers=false:.  api/grpc/EventService.proto
