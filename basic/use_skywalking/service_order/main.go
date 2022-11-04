package main

import (
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"swwgo/basic/use_skywalking/conf"
	"time"

	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
)

const (

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

func CalcAuth(c *gin.Context, userToken string) {
	req, err := http.NewRequest(http.MethodGet, conf.ServiceAuthUrl + "?token=" + userToken, nil)
	DealErr(err)

	// 创建span
	resSpan, err := tracer.CreateExitSpan(c.Request.Context(), "invoke -"+ conf.ServiceAuth, "xxxxx", func(headerKey, headerValue string) error {
		req.Header.Set(headerKey, headerValue)
		return nil
	})
	DealErr(err)
	defer func() {
		resSpan.End()
	}()

	// 调用第三方http接口
	fmt.Printf("调用： %s\n", userToken)
	resp, err := http.DefaultClient.Do(req)
	DealErr(err)
	all, err := ioutil.ReadAll(resp.Body)
	DealErr(err)
	fmt.Printf("接受： %s\n", all)
}

func main() {
	startSkyWalking()

	r := gin.Default()

	// gin 使用 sky 自带的 middleware
	r.Use(v3.Middleware(r, tracer))

	r.GET("/order_pay", func(c *gin.Context) {
		println("order_pay 业务处理111...")

		println("order_pay 调用auth服务...")
		CalcAuth(c, "sww的xxxxToken")

		println("order_pay 订单处理...")
		OrderPay(c,"songweiwei", "199112312912931")

		c.JSON(200, gin.H{
			"message" : "ping",
		})
	})

	r.Run(":80")
}
