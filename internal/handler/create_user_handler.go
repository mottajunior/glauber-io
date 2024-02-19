package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mottajunior/glauber-io/internal/config"
	"github.com/mottajunior/glauber-io/internal/repository"
	usecase "github.com/mottajunior/glauber-io/internal/usecases"
)

func CreateUser(ctx *gin.Context) {
	repository := repository.NewUserDatabaseRepository(config.GetDatabaseConnection())
	usecase := usecase.NewCreateUserUseCase(repository)
	usecase.Execute(ctx)
}
