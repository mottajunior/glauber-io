package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mottajunior/glauber-io/internal/routes"
)

func main() {
	server := gin.Default()
	routes.ConfigureRoutes(server)
	server.Run()
}
