package main

import (
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"

	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
)

const (
	serverName = "service-111"
	skyAddr = "192.168.71.131:11800"

	service222Url = "http://127.0.0.1:81/hello222"
	service222Name = "service-222"
)

var tracer *go2sky.Tracer

func startSkyWalking() {
	var rp go2sky.Reporter
	var err error
	// skyAddr 是 skywaling 的 grpc 地址，默认是 localhost:11800， 默认心跳检测时间是 1s
	rp, err = reporter.NewGRPCReporter(skyAddr, reporter.WithCheckInterval(5*time.Second))
	if err != nil {
		panic(err)
	}

	// 初始化一个tracer， 一个服务只需要一个tracer，其含义是这个服务名称
	tracer, err = go2sky.NewTracer(serverName, go2sky.WithReporter(rp))
	if err != nil {
		panic(err)
	}
}

func DealErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	startSkyWalking()

	r := gin.Default()

	// gin 使用 sky 自带的 middleware
	r.Use(v3.Middleware(r, tracer))

	r.GET("/hello1", func(c *gin.Context) {
		println("hello1 业务处理...")

		//// 调用第三方http接口
		req, err := http.NewRequest(http.MethodGet, service222Url, nil)
		DealErr(err)

		//resSpan, err := tracer.CreateExitSpan(c.Request.Context(), "invoke -"+ service222Name, "perr", func(headerKey, headerValue string) error {
		//	req.Header.Set(headerKey, headerValue)
		//	return nil
		//})
		//DealErr(err)
		//
		//resSpan.Log(time.Now(), "开始请求")
		//
		resp, err := http.DefaultClient.Do(req)
		DealErr(err)
		all, err := ioutil.ReadAll(resp.Body)
		DealErr(err)
		fmt.Printf("接受消息： %s\n", all)

		//resSpan.Tag(go2sky.TagHTTPMethod, http.MethodGet)
		//resSpan.Tag(go2sky.TagURL, service222Url)
		//resSpan.Log(time.Now(), "结束请求")
		//
		//resSpan.End()

		c.JSON(200, gin.H{
			"message" : "ping",
		})

		//// LocalSpan 可以理解为本地日志的 tracer，一般用户当前应用
		//span, ctx, err := tracer.CreateLocalSpan(c.Request.Context())
		//log.Println(err)
		//// 每一个 span 都有一个名字去标实操作的名称！
		//span.SetOperationName("UserInfo")
		//// 记住重新设置一个 ctx，再其次这个 ctx 不是 gin 的 ctx，而是 http request 的 ctx
		//c.Request = c.Request.WithContext(ctx)
		//
		//span.Log(time.Now(), "12312312")
		//
		//span.End()
	})

	r.Run(":80")
}
