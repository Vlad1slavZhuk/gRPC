package main

import (
	api "gRPC/api/protoc"
	s "gRPC/internal/pkg/storage/grpc"
	mem "gRPC/internal/pkg/storage/in-memory"
	"log"
	"net"

	"google.golang.org/grpc"
)

// var port string

// func init() {
// 	if port = os.Getenv("GRPC_PORT"); port == "" {
// 		log.Fatal("Set GRPC_PORT!")
// 	}
// 	port = ":" + port
// }

func main() {
	ls, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	// service := os.Getenv("STORAGE_SERVICE")
	// switch service {
	// case "", "memory":
	// 	log.Printf("GRPC server: launching on port %v with in-memory storage\n", port)
	// 	api.RegisterSliStorageServiceServer(srv, grpcsrv.NewGRPCStorage(mem.New()))
	// case "redis":
	// 	log.Printf("GRPC server: launching on port %v with redis storage\n", port)
	// 	api.RegisterSliStorageServiceServer(srv, grpcsrv.NewGRPCStorage(redis.New()))
	// case "postgres":
	// 	log.Printf("GRPC server: launching on port %v with postgres storage\n", port)
	// 	api.RegisterSliStorageServiceServer(srv, grpcsrv.NewGRPCStorage(postgres.New()))
	// case "mongo":
	// 	log.Printf("GRPC server: launching on port %v with mongodb storage\n", port)
	// 	api.RegisterSliStorageServiceServer(srv, grpcsrv.NewGRPCStorage(mongo.New()))
	// default:
	// 	log.Fatal("Set valid STORAGE_SERVICE!")
	// }
	api.RegisterServiceProtobufServer(srv, s.NewStorageGrpcServer(mem.NewStorage()))
	log.Println("Start grpc on localhost:9000")
	log.Fatal(srv.Serve(ls))
}
