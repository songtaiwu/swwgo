package main

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
)

const (
	serverName = "service-222"
	skyAddr = "192.168.71.131:11800"
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

func main() {
	startSkyWalking()

	r := gin.Default()

	// skyAddr 是 skywaling 的 grpc 地址，默认是 localhost:11800， 默认心跳检测时间是 1s
	rp, err := reporter.NewGRPCReporter(skyAddr, reporter.WithCheckInterval(5*time.Second))
	log.Println(err)

	// 初始化一个tracer， 一个服务只需要一个tracer，其含义是这个服务名称
	tracer, err := go2sky.NewTracer(serverName, go2sky.WithReporter(rp))
	log.Println(err)

	// gin 使用 sky 自带的 middleware
	r.Use(v3.Middleware(r, tracer))

	r.GET("/hello222", func(c *gin.Context) {
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

	r.Run(":81")
}
