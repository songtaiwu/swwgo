package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"swwgo/use_gokit/v2/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial(":1001", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("did not connect")
	}
	defer conn.Close()

	c := proto.NewArticleClient(conn)

	// 可以传递一些meta信息
	paramCtx := context.Background()
	pairs := metadata.Pairs("k1", "v1", "k2", "v2")
	outgoingContext := metadata.NewOutgoingContext(paramCtx, pairs)

	ctx, cancel := context.WithTimeout(outgoingContext, time.Second)
	defer cancel()
	res, err := c.ArticleAdd(ctx, &proto.AddRequest{Name: "xiaoming", Content: "beijing"})
	if err != nil {
		log.Fatalf("could not add article %v", err)
	}
	log.Printf("response: %s", res.Data)
}

