package routes

import (
	handlers "glauber-io/internal/handlers/account"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(server *gin.Engine) {
	group := server.Group("api/v1")

	group.GET("account/create", handlers.CreateAccount)
	group.POST("account/login", handlers.Login)
	group.POST("account/forgot-password", handlers.ForgotPassword)
}
