package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "url_shortener/proto"
	editor "url_shortener/server/editorUrl"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("fatal error listening tcp: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEditorUrlServer(grpcServer, &editor.GRPCServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fatal error: %v", err)
	}
}
