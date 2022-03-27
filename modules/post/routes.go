package post

import "github.com/gin-gonic/gin"

func InitRoutes(route *gin.Engine) {

	postsGroup := route.Group("api/v1")
	{
		postsGroup.GET("/posts", GetPosts)
	}

}
