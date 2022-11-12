package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"swwgo/basic/use_grpc/v1/proto"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:1001")
	if err != nil {
		log.Fatalln("failed to listen")
	}

	grpcServer := grpc.NewServer()
	proto.RegisterArticleServer(grpcServer, new(ArticleServer))
	grpcServer.Serve(lis)

	println(123)
}

type ArticleServer struct {
	proto.UnimplementedArticleServer
}

func (s ArticleServer) ArticleAdd(ctx context.Context, request *proto.AddRequest) (*proto.AddResponse, error) {
	res := &proto.AddResponse{
		Data: request.Name + " " + request.Content,
	}
	return res, nil
}