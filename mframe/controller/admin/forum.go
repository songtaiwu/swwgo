package admin

import (
	"github.com/gin-gonic/gin"
	"swwgo/mframe/package/logging"
	"swwgo/mframe/service"
)

var userService = new(service.UserService)

func ForumGet(c *gin.Context) {

	logging.Debug("asdasdasd")
	c.JSON(200, gin.H{
		"code" : 0,
		"msg" : "success",
		"data" : "admin ForumGet",
	})
}

func ForumList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code" : 0,
		"msg" : "success",
		"data" : "admin ForumList",
	})
}