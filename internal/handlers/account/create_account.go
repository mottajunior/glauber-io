package handlers

import "github.com/gin-gonic/gin"

func CreateAccount(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"isabela": "teamo",
	})
}
