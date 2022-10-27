package router

import (
	"github.com/gin-gonic/gin"
	"swwgo/mframe/controller/front"
)

func SetFrontRouterGroup(g *gin.RouterGroup) {
	forum := g.Group("forum")
	{
		forum.GET("/get", front.ForumGet)
		forum.GET("/list", front.ForumList)
	}
}
