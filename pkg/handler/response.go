package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// создаём структуру для ошибок
type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	// уведомляем в консоли об ошибке
	logrus.Error(message)
	//AbortWithStatusJSON прерывает все последующие обратботчики и отправляет ответ с статус кодом и телом запроса
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
