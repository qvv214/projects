package editorUrl

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"time"

	pb "url_shortener/proto"
	db "url_shortener/store"
)

const symbols = "0123456789_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type GRPCServer struct{}

func (s *GRPCServer) Create(ctx context.Context, req *pb.LongUrl) (res *pb.ShortUrl, err error) {
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

	for i := range b {
		b[i] = symbols[generator.Intn(len(symbols))]
	}

	connection := new(db.Store)

	if err := connection.Open(); err != nil {
		log.Fatalf("fatal open db: %v", err)
	}

	connection.AddUrl(req.Url, string(b))

	defer connection.Close()

	return &pb.ShortUrl{Url: string(b)}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *pb.ShortUrl) (res *pb.LongUrl, err error) {
	fmt.Printf("short Url: %s", &pb.ShortUrl{})
	return &pb.LongUrl{}, nil
}
