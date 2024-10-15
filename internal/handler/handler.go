package handler

import (
	"Afisha/internal/handler/extension"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *extension.Services
}

func NewHandler(service *extension.Services) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/auth")
	{
		api.POST("/signIn", h.SignIn)
		api.POST("/signUp", h.SignUp)
		api.POST("/signOut", h.SignOut)
	}

	return router
}
