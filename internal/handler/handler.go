package handler

import (
	"EffectiveMobile/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "EffectiveMobile/docs"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	music := router.Group("/music")
	{
		music.POST("/", h.createSong)
		music.POST("/:id", h.createVerse)
		music.POST("/info", h.createInfo)
		music.GET("/", h.getLibrary)
		music.GET("/:id", h.getSong)
		music.PUT("/", h.updateSong)
		music.DELETE("/:id", h.deleteSong)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
