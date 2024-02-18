package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mottajunior/glauber-io/internal/handler"
)

func configureRoutes(server *gin.Engine) {
	group := server.Group("api/v1")
	{
		group.POST("account/create", handler.CreateAccount)
		group.POST("account/login", handler.Login)
		group.POST("account/forgot-password", handler.ForgotPassword)
	}
}
