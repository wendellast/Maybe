package handler

import "github.com/gin-gonic/gin"

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": " GET Tudo Okay!",
	})
}

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "POST Tudo Okay!",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "DEL Tudo Okay!",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "PUT Tudo Okay!",
	})
}

func ListOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "GEt Tudo Okay!",
	})
}
