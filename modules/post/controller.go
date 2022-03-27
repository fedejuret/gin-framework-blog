package post

import "github.com/gin-gonic/gin"

func GetPosts(c *gin.Context) {

	posts := GetAllPosts()

	c.JSON(200, gin.H{
		"data": posts,
	})
}
