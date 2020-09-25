package main

import (
	"fmt"
	"github.com/serkancetintas/grpc/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

type testServer struct {
	chat.UnimplementedChatServiceServer
}

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &testServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
