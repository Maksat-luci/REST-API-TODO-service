package handler

import (
	"net/http"

	todo "github.com/Maksat-luci/REST-API-TODO-service"
	"github.com/gin-gonic/gin"
)

// signUp метод который валидирует данные а далее отправляет в бизнес логику, далее тот использует слой
// базы данных, а база данных возвращает id который принимает хендлер через бизнес логику
func (h *Handler) signUp(c *gin.Context) {
	// структура юзера для записи данных нового юзера
	var input todo.User
	// записываем данные в нашу структуру который одновременно её и валидирует
	if err := c.BindJSON(&input); err != nil {
		// уведомляем в консоли и в ответе
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// переходим в слой бизнес логики и передаём ей наши данные
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		// если произошла ошибка то уведомляем
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	// если всё ок то с помощью gin.Context.Json отправляем ответ в виде json
	// первым аргументом статус ответа вторым в качестве ключ, значения, мапу так как json устроен по принципу ключа значения
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	// структура signInInput для записи данных зарегистрированного юзера
	var input signInInput
	// записываем данные в нашу структуру который одновременно её и валидирует
	if err := c.BindJSON(&input); err != nil {
		// уведомляем в консоли и в ответе
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	// если всё ок то с помощью gin.Context.Json отправляем ответ в виде json
	// первым аргументом статус ответа вторым в качестве ключ, значения, мапу так как json устроен по принципу ключа значения
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
