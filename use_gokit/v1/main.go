package main

import (
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"swwgo/use_gokit/v1/app/my_endpoint"
	"swwgo/use_gokit/v1/app/my_service"
	"swwgo/use_gokit/v1/app/my_tansport"
)

func main() {
	// 1.先创建定义的业务处理接口实现， service
	svc := my_service.NewService()

	// 2.再创建业务服务的函数签名，endpoint
	endpointArticleAdd := my_endpoint.MakeEndPointArticleAdd(svc)

	// 3.使用kit创建handler
	handler := httpTransport.NewServer(
		endpointArticleAdd,
		my_tansport.DecodeArticleAddRequest,
		my_tansport.EncodeArticleAddResponse,
	)

	// 4. http路由和服务
	m := http.NewServeMux()
	m.Handle("/article", handler)
	fmt.Println("server run 0.0.0.0:8888")
	err := http.ListenAndServe("0.0.0.0:8888", m)
	if err != nil {
		println(err.Error())
	}
	select {

	}
}
