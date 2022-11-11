package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func OrderPay(c *gin.Context, username string, orderId string) string {
	// LocalSpan 可以理解为本地日志的 tracer，一般用户当前应用
	span, _, _ := tracer.CreateLocalSpan(c.Request.Context())
	defer func() {
		span.End()
	}()

	// 每一个 span 都有一个名字去标实操作的名称！
	span.SetOperationName("OrderPay")
	span.Tag("param_username", username)
	span.Tag("param_orderid", orderId)

	// 记住重新设置一个 ctx，再其次这个 ctx 不是 gin 的 ctx，而是 http request 的 ctx
	c.Request = c.Request.WithContext(c)
	span.Log(time.Now(), "12312312")

	ret := "1998111111111"
	span.Tag("return", ret)
	return ret
}
