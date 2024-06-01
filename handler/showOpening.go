package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wendellast/Maybe/schema"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	SendSuccess(ctx, "show-opening", opening)

}
