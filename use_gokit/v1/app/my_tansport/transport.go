package my_tansport

import (
	"context"
	"errors"
	"github.com/goccy/go-json"
	"net/http"
	"swwgo/use_gokit/v1/app/my_service"
)

// Transport 负责HTTP、gRPC、thrift等相关协议的请求逻辑

// 对每一个请求都要实现一对参数解码和返回值编码的函数签名。
// DecodeRequest & EncodeResponse 函数签名是固定的。
// func DecodeRequest(c context.Context, request *http.Request) (interface{}, error)
// func EncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error

// DecodeRequest解码，请求参数封装为Endpoint中定义的Request格式
func DecodeArticleAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var param my_service.ArticleAddRequest
	param.Name = r.FormValue("name")
	param.Content = r.FormValue("content")

	if param.Name == "" {
		return nil, errors.New("name不能为空")
	}

	return param, nil
}

// EncodeResponse编码，把业务的响应封装成想要的结构
func EncodeArticleAddResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// 这里将Response返回成有效的json格式给http
	// 设置请求头
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 使用内置json包转换
	return json.NewEncoder(w).Encode(response)
}