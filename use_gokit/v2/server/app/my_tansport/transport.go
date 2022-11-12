package my_tansport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"swwgo/use_gokit/v2/proto"
)

type grpcServer struct {
	addArticle kitgrpc.Handler
	proto.UnimplementedArticleServer
}

func (g grpcServer) ArticleAdd(ctx context.Context, request *proto.AddRequest) (*proto.AddResponse, error) {
	_, resp, err := g.addArticle.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.AddResponse), nil
}

func NewArticleServer(ctx context.Context, endpoint endpoint.Endpoint) proto.ArticleServer {
	return &grpcServer{
		addArticle: kitgrpc.NewServer(
			endpoint,
			DecodeArticleAddRequest,
			EncodeArticleAddResponse,
			),
	}
}
