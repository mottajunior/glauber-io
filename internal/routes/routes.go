package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/mottajunior/glauber-io/internal/handlers/account"
)

func ConfigureRoutes(server *gin.Engine) {
	group := server.Group("api/v1")
	{
		group.GET("account/create", handlers.CreateAccount)
		group.POST("account/login", handlers.Login)
		group.POST("account/forgot-password", handlers.ForgotPassword)
	}
}
