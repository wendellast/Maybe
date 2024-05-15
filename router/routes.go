package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(route *gin.Engine) {
	v1 := route.Group("/api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"msg": "Tudo Okay!",
			})
		})
	}
}
