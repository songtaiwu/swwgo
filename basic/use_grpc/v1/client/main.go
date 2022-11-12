package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"swwgo/basic/use_grpc/v1/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial(":1001", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("did not connect")
	}
	defer conn.Close()

	c := proto.NewArticleClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.ArticleAdd(ctx, &proto.AddRequest{Name: "xiaoming", Content: "beijing"})
	if err != nil {
		log.Fatalf("could not add article %v", err)
	}
	log.Printf("response: %s", res.Data)
}

