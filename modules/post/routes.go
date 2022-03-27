package post

import "github.com/gin-gonic/gin"

func InitRoutes(route *gin.Engine) {

	postsGroup := route.Group("api/v1")
	{
		postsGroup.GET("/posts/:id", GetPostController)
		postsGroup.GET("/posts", GetPostsController)

		postsGroup.POST("/posts", CreatePostController)
	}

}
