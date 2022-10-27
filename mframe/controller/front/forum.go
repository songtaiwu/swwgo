package front

import "github.com/gin-gonic/gin"


func ForumGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"code" : 0,
		"msg" : "success",
		"data" : "front ForumGet",
	})
}

func ForumList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code" : 0,
		"msg" : "success",
		"data" : "front ForumList",
	})
}