package handler

import (
	"Afisha/internal/handler/extension"
	"Afisha/pkg/tokenManager"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *extension.Services
	manager tokenManager.ITokenManager
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
	events := router.Group("/events")
	{
		events.GET("/:id", h.GetEventById)
		events.GET("/all-events", h.AllEventsPage)
		events.GET("/get-by-name", h.GetEventsByName)
		events.GET("/get-by-filters", h.GetEventsByFilters)
		events.POST("/expectations", h.UserExpected)
		events.POST("/buy-tickets", h.BuyTickets)
		events.POST("/create-event", h.CreateEvent)
		events.DELETE("delete-event", h.DeleteEvent)
	}
	return router
}
