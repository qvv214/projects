package main

import (
	"context"
	"flag"
	"log"
	client "url_shortener/proto"

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
			c := client.NewEditorUrlClient(conn)
			c.Create(context.Background(), &client.LongUrl{Url: flag.Arg(1)})
		}
	case "Get":
		{
			c := client.NewEditorUrlClient(conn)
			c.Get(context.Background(), &client.ShortUrl{Url: flag.Arg(1)})
		}
	default:
		{
			log.Fatal("not function")
		}
	}

}
