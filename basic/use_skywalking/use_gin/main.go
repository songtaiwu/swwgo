package main

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"time"

	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
)

const (
	serverName = "sww-demo1"
	skyAddr = "127.0.0.1:11800"
)

func main() {
	r := gin.Default()


	// skyAddr 是 skywaling 的 grpc 地址，默认是 localhost:11800， 默认心跳检测时间是 1s
	rp, err := reporter.NewGRPCReporter(skyAddr, reporter.WithCheckInterval(5*time.Second))
	panic(err)

	// 初始化一个tracer， 一个服务只需要一个tracer，其含义是这个服务名称
	tracer, err := go2sky.NewTracer(serverName, go2sky.WithReporter(rp))
	panic(err)

	// gin 使用 sky 自带的 middleware
	r.Use(v3.Middleware(r, tracer))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "ping",
		})

		// LocalSpan 可以理解为本地日志的 tracer，一般用户当前应用
		span, ctx, err := tracer.CreateLocalSpan(c.Request.Context())
		panic(err)

		// 每一个 span 都有一个名字去标实操作的名称！
		span.SetOperationName("UserInfo")

		c.Request = c.Request.WithContext(ctx)

	})

	r.Run(":80")
}
