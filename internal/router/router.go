package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	server := gin.Default()
	configureRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run("0.0.0.0:" + port)
}
