package controller

import (
	"abix360/src/usecase"
	"abix360/src/view/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePassword(c *gin.Context) {
	var req dto.CreatePasswordDTO
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	useCase := usecase.CreatePasswordUseCase{}
	passwordDto, err := useCase.Execute(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusAccepted, err.Error())
		return
	}
	c.JSON(http.StatusOK, passwordDto)

}
