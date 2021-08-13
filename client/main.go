package main

import (
	"context"
	"flag"
	"log"
	pb "url_shortener/proto"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fatal err: %v", err)
	}

	switch flag.Arg(0) {
	case "Create":
		{
			c := pb.NewEditorUrlClient(conn)
			c.Create(context.Background(), &pb.LongUrl{Url: flag.Arg(1)})
		}
	case "Get":
		{
			c := pb.NewEditorUrlClient(conn)
			c.Get(context.Background(), &pb.ShortUrl{Url: flag.Arg(1)})
		}
	default:
		{
			log.Fatal("not function")
		}
	}

}
