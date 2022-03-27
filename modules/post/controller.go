package post

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPostsController(c *gin.Context) {

	posts := GetAllPosts()

	c.JSON(200, gin.H{
		"data": posts,
	})
}

func GetPostController(c *gin.Context) {

	id := c.Param("id")

	postId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid post id",
		})
		return
	}

	post := GetPostById(postId)

	c.JSON(200, gin.H{
		"data": post,
	})
}

func CreatePostController(c *gin.Context) {

	var post Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	post = CreatePost(post)

	c.JSON(200, gin.H{
		"data": post,
	})
}
