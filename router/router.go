package router

import "github.com/gin-gonic/gin"

func Initializer() {
	router := gin.Default()
	InitializeRoutes(router)

	router.Run(":8080")
}
