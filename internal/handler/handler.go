package handler

import (
	"Afisha/internal/handler/extension"
	"Afisha/pkg/tokenManager"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

type Handler struct {
	service *extension.Services
	manager tokenManager.ITokenManager
}

func NewHandler(service *extension.Services) *Handler {
	signingKey := os.Getenv("SIGNING_KEY")
	return &Handler{
		service: service,
		manager: tokenManager.NewTokenManager(signingKey),
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
	events := router.Group("/events")
	{
		events.GET("/:id", h.GetEventById)
		events.POST("/all-events", h.AllEventsPage)
		events.POST("/popular", h.PopularPage)
		events.POST("/expectations", h.UserExpected)
	}
	return router
}
