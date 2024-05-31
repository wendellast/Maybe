package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wendellast/Maybe/schema"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validade(); err != nil {
		logger.Errof("Validation Error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := schema.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errof("Erro creating opening: %v ", err.Error())
		SendError(ctx, http.StatusInternalServerError, "Erro creating opening database")
		return

	}

	SendSuccess(ctx, "create-opening", opening)

}
