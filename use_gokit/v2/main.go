package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"swwgo/use_gokit/v2/proto"
	"swwgo/use_gokit/v2/server/app/my_endpoint"
	"swwgo/use_gokit/v2/server/app/my_service"
	"swwgo/use_gokit/v2/server/app/my_tansport"
)

func main() {
	// 1.先创建定义的业务处理接口实现， service
	svc := my_service.NewService()

	// 2.再创建业务服务的函数签名，endpoint
	endpointArticleAdd := my_endpoint.MakeEndPointArticleAdd(svc)

	ctx := context.Background()
	handler := my_tansport.NewArticleServer(ctx, endpointArticleAdd)

	lis, err := net.Listen("tcp", "0.0.0.0:1001")
	if err != nil {
		log.Fatalln("failed to listen")
	}

	grpcServer := grpc.NewServer()
	proto.RegisterArticleServer(grpcServer, handler)
	grpcServer.Serve(lis)
}
