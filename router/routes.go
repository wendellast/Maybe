package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wendellast/Maybe/handler"
)

func InitializeRoutes(route *gin.Engine) {

	handler.InitializerHandler()
	v1 := route.Group("/api/v1")
	{
		v1.GET("/opening/show/", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/opening", handler.ListOpeningHandler)

	}
}
