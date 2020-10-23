package main

import (
	"gRPC/internal/pkg/server"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := server.NewServer()
	s.SetHandlers()
	s.SetStorage()
	//server.SetConfig()
	s.Run()
}
