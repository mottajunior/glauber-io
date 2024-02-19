package usecase

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/mottajunior/glauber-io/internal/dtos"
	"github.com/mottajunior/glauber-io/internal/repository"
	"github.com/mottajunior/glauber-io/internal/schemas"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUseCase(repo repository.UserRepository) *CreateUserUseCase {
	usecase := CreateUserUseCase{userRepository: repo}
	return &usecase
}

func (usecase CreateUserUseCase) Execute(ctx *gin.Context) {
	request := dto.CreateUserRequest{}
	ctx.BindJSON(&request)

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

	if err := usecase.userRepository.Save(user); err != nil {
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
