package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"swwgo/basic/use_jaeger/v1/jaeger"
)

func main() {
	jaeger.StartJaeger()

	r := gin.New()

	r.Use(jaeger.Middleware())

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code" : 0,
			"msg" : "hello tom",
		})
	})

	s := &http.Server{
		Addr: ":80",
		Handler: r,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Panicf("start http failed: %v", err)
	}
}


