package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/mottajunior/glauber-io/internal/dto"
	"github.com/mottajunior/glauber-io/internal/repository"
	"github.com/mottajunior/glauber-io/internal/schemas"
)

type UserHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	uc := UserHandler{userRepository: repo}
	return &uc
}

func (usrHandler UserHandler) Create(ctx *gin.Context) {
	request := dto.CreateUserRequest{}
	_ = ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		slog.Error("Error when validate request: ", err)
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Name:  request.Name,
		Email: request.Email,
		Age:   request.Age,
	}

	if err := usrHandler.userRepository.Save(user); err != nil {
		slog.Error("Error when save user on database: ", err)
		SendError(ctx, http.StatusInternalServerError, "error creating user on database")
		return
	}

	response := dto.CreateUserResponse{
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}

	SendSuccess(ctx, "create user", response)
}
