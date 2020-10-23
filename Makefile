go:
	go run cmd/http/main.go

protoc:
	protoc -I api/ api/protobuf.proto --go_out=plugins=grpc:api