package handler

import (
	"avito-segment/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		segments := api.Group("/segments")
		{
			segments.POST("/", h.createSegment)
			segments.DELETE("/:slug", h.deleteSegment)
		}

		users := api.Group("/users")
		{
			users.GET("/:user_id/segments", h.getUserSegments)
			users.POST("/:user_id/segments", h.updateUserSegments)
		}

	}

	return router
}
