package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/url"
	"time"
	srv "url_shortener/proto"

	"google.golang.org/grpc"
)

type GRPCServer struct{}

const symbols = "0123456789_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func (s *GRPCServer) Create(ctx context.Context, req *srv.LongUrl) (res *srv.ShortUrl, err error) {

	fmt.Println("server work!")

	ur, err := url.ParseRequestURI(req.Url)
	if err != nil {
		log.Fatalf("fatal error: %v", err)
	}

	u, err := ur.Parse(req.Url)
	if err != nil || u.Host == "" {
		log.Fatalf("fatal error: %v", err)
	}

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())

	for i := range b {
		b[i] = symbols[generator.Intn(len(symbols))]
	}

	return &srv.ShortUrl{Url: string(b)}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *srv.ShortUrl) (res *srv.LongUrl, err error) {
	return &srv.LongUrl{Url: "test"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("fatal error: %v", err)
	}

	grpcServer := grpc.NewServer()
	srv.RegisterEditorUrlServer(grpcServer, &GRPCServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fatal error: %v", err)
	}
}
