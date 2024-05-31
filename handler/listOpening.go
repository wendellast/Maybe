package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wendellast/Maybe/schema"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schema.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "Erro listing opening")
		return
	}

	SendSuccess(ctx, "List-Opening", openings)
}
