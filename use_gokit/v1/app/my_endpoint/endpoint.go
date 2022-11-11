package my_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"swwgo/use_gokit/v1/app/my_service"
)

// MakeEndPointArticleAdd 创建关于业务的构造函数
// 传入service层定义的相关业务接口
// 返回 endpoint.Endpoint， 实际就是一个函数签名
func MakeEndPointArticleAdd(svc my_service.IService) endpoint.Endpoint {
	// 这里使用闭包，可以在这里做一些业务的处理
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// request是对应请求过来时传入的参数，实际上是Transport中一个decode函数处理得到的
		// 需要进行下断言
		req := request.(my_service.ArticleAddRequest)

		// 这里就是调用service层定义的业务逻辑
		// 把拿到的数据作为参数
		res := svc.ArticleAdd(ctx, req)

		// 返回值可以是任意的，不过根据规范要返回我们刚才定义好的返回对象
		return res, nil
	}
}
