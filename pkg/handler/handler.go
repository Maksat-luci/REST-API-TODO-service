package handler

import (
	"github.com/Maksat-luci/REST-API-TODO-service/pkg/service"
	"github.com/gin-gonic/gin"
)

//Handler структура которая содержит в себе интерфейс нижнего слоя service
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

//InitRoutes иницилиризует пути тоесть даёт понять функции listenAndServe какие пути нужно слушать
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up",h.signUp)
		auth.POST("/sign-in",h.signIn)
	}

	api := router.Group("/api",h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/",h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id",h.getListByID)
			lists.PUT("/:id",h.updateList)
			lists.DELETE("/:id",h.deleteList)

			items := lists.Group("id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemByID)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
