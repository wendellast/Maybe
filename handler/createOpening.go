package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wendellast/Maybe/schema"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
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
