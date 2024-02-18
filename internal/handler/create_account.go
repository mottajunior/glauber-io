package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mottajunior/glauber-io/internal/schemas"
)

func CreateAccount(ctx *gin.Context) {
	request := CreateAccountRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		slog.Error("Error when validate request: ", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	account := schemas.Account{
		Name:  request.Name,
		Email: request.Email,
		Age:   request.Age,
	}

	if err := db.Create(&account).Error; err != nil {
		slog.Error("Error when save account in database: ", err)
		sendError(ctx, http.StatusInternalServerError, "error creating account on database")
		return
	}

	sendSuccess(ctx, "create account", account)
}
