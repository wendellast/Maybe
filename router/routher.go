package router

import "github.com/gin-gonic/gin"

func Initializer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message:": "Pong",
		})

	})

	r.Run()
}
