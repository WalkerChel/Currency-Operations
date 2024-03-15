package handler

import (
	"currency-operations/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	ApiKey   string
}

func NewHandler(services *service.Service, apiKey string) *Handler {
	return &Handler{
		services: services,
		ApiKey:   apiKey,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		accounts := api.Group("/accounts")
		{
			accounts.POST("/", h.CreateAccount)
			accounts.GET("/", h.GetAllAccounts)
			accounts.GET("/:id", h.GetAccountById)
			accounts.PUT("/deposit/:id", h.Deposit)
		}
	}

	return router
}
