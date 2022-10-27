package router

import (
	"github.com/gin-gonic/gin"
	"swwgo/mframe/controller/admin"
)

func SetAdminRouterGroup(g *gin.RouterGroup) {
	forum := g.Group("forum")
	{
		forum.GET("/get", admin.ForumGet)
		forum.GET("/list", admin.ForumList)
	}
}
