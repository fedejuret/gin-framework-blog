package main

import (
	"github.com/gin-gonic/gin"
	"go.fedejuret.blog/modules/post"
)

func main() {

	routes := gin.Default()

	post.Init(routes)

	routes.Run()

}
