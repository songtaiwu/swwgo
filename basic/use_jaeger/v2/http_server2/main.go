package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	jaeger2 "swwgo/basic/use_jaeger/v2/jaeger"
)

func main() {
	jaeger2.StartJaeger("http_server2")

	r := gin.New()

	r.Use(jaeger2.Middleware())

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code" : 0,
			"msg" : "I am http-server2",
		})
	})

	s := &http.Server{
		Addr: ":8002",
		Handler: r,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Panicf("start http failed: %v", err)
	}
}


