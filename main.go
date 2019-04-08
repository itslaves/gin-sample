package main

import (
	"gin-sample/article"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/articles", article.Retrieve)
	r.POST("/articles", article.Create)
	r.PUT("/articles/:id", article.Update)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}