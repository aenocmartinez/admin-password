package controller

import (
	"abix360/src/usecase"
	"abix360/src/view/dto"
	formrequest "abix360/src/view/form-request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePassword(c *gin.Context) {
	var req formrequest.CreatePasswordFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	useCase := usecase.CreatePasswordUseCase{}
	passwordDto, err := useCase.Execute(dto.PasswordDTO{
		Name:     req.Name,
		User:     req.Url,
		Password: req.Password,
		Url:      req.Url,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusAccepted, err.Error())
		return
	}
	c.JSON(http.StatusOK, passwordDto)
}

func PasswordList(c *gin.Context) {
	useCase := usecase.PasswordListUseCase{}
	passwordList, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, passwordList)
}

func FindPasswordById(c *gin.Context) {
	strId := c.Query("id")
	if len(strId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.FindPasswordUseCase{}
	password, err := useCase.Execute(int64(id))
	if err != nil {
		c.AbortWithStatusJSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, password)
}

func UpdatePassword(c *gin.Context) {
	var req formrequest.UpdatePasswordDTOFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	useCase := usecase.UpdatePasswordUseCase{}
	passwordDto, err := useCase.Execute(dto.PasswordDTO{
		Id:       req.Id,
		Name:     req.Name,
		User:     req.Url,
		Password: req.Password,
		Url:      req.Url,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusAccepted, err.Error())
		return
	}
	c.JSON(http.StatusOK, passwordDto)
}

func DeletePassword(c *gin.Context) {
	strId := c.Query("id")
	if len(strId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.DeletePasswordUseCase{}
	err = useCase.Execute(int64(id))
	if err != nil {
		c.AbortWithStatusJSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
