package my_service

import (
	"context"
)

// IService 用于定义业务方法的接口
type IService interface {
	ArticleAdd(ctx context.Context, param ArticleAddRequest) ArticleAddResponse
}

// baseService 用于实现上面定义的接口
type baseService struct {
	// 根据业务需求填充结构体...
}

func (b baseService) ArticleAdd(ctx context.Context, param ArticleAddRequest) ArticleAddResponse {
	return ArticleAddResponse{
		Data: param.Name + param.Content,
	}
}

func NewService() IService {
	return &baseService{}
}
