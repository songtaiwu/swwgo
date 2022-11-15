package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"io/ioutil"
	"log"
	"net/http"
	jaeger2 "swwgo/basic/use_jaeger/v2/jaeger"
	"time"
)

func main() {
	jaeger2.StartJaeger("http_server1")

	r := gin.New()

	r.Use(jaeger2.Middleware())

	r.GET("/hello", func(c *gin.Context) {
		time.Sleep(time.Second * 1)
		// 再创建一个span
		spanC, _ := c.Get(jaeger2.SpanCTX)
		spanCtx := spanC.(context.Context)
		subSpan, _ := opentracing.StartSpanFromContext(spanCtx, "asdsad")
		defer subSpan.Finish()
		subSpan.SetBaggageItem("good_key1", "good_value1")
		subSpan.SetBaggageItem("good_key2", "good_value2")

		// 构建GET请求
		client := &http.Client{}
		request, _ := http.NewRequest("GET", "http://127.0.0.1:8002/hello", nil)

		// 把trace信息传到下一个http服务
		opentracing.GlobalTracer().Inject(subSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(request.Header))

		resp, err := client.Do(request)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return
		}

		c.JSON(200, gin.H{
			"code" : 0,
			"msg" : string(body),
		})
	})

	s := &http.Server{
		Addr: ":8001",
		Handler: r,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Panicf("start http failed: %v", err)
	}
}


