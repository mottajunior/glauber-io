package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mottajunior/glauber-io/internal/handler"
)

func configureRoutes(server *gin.Engine) {
	group := server.Group("api/v1")
	{
		group.POST("user", handler.CreateUser)
	}
}
