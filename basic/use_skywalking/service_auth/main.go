package main

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"swwgo/basic/use_skywalking/conf"
	"time"

	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
)

var tracer *go2sky.Tracer

func startSkyWalking() {
	var rp go2sky.Reporter
	var err error
	// skyAddr 是 skywaling 的 grpc 地址，默认是 localhost:11800， 默认心跳检测时间是 1s
	rp, err = reporter.NewGRPCReporter(conf.SkyAddr, reporter.WithCheckInterval(5*time.Second))
	if err != nil {
		panic(err)
	}

	// 初始化一个tracer， 一个服务只需要一个tracer，其含义是这个服务名称
	tracer, err = go2sky.NewTracer(conf.ServiceAuth, go2sky.WithReporter(rp))
	if err != nil {
		panic(err)
	}
}

func main() {
	startSkyWalking()

	r := gin.Default()

	// gin 使用 sky 自带的 middleware
	r.Use(v3.Middleware(r, tracer))

	r.GET("/auth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "ping",
		})
	})

	r.Run(":81")
}
