package main

import (
	"glauber-io/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.ConfigureRoutes(server)
	server.Run()
}
