package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()


	// 前台接口
	front := r.Group("api/front")
	SetFrontRouterGroup(front)

	// 后台接口
	admin := r.Group("api/admin")
	SetAdminRouterGroup(admin)


	return r
}