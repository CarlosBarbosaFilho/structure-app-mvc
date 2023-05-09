package service

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/config"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/helpers"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/models"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/repositories"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/request"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/response"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	IUserRepository repositories.IUserRepository
	Validate        *validator.Validate
}

func loggerMethods() *config.Logger {
	logger := config.NewLogger("Logging - Routes Configuration")
	return logger
}

func NewUserServiceImpl(repository repositories.IUserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{IUserRepository: repository, Validate: validate}
}

// CreateUser implements UserService
func (service *UserServiceImpl) CreateUser(user request.UserRequest) {
	loggerMethods().Info("Service:: Creating user to save")
	err := service.Validate.Struct(user)
	helpers.ErrorPanic(err)

	userSave := models.Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	service.IUserRepository.CreateUser(userSave)
	loggerMethods().Info("Service:: User save with success")
}

// DeleteUser implements UserService
func (*UserServiceImpl) DeleteUser(id uint) {
	panic("unimplemented")
}

// GetUserById implements UserService
func (*UserServiceImpl) GetUserById(id uint) response.UserResponse {
	panic("unimplemented")
}

// ListUsers implements UserService
func (*UserServiceImpl) ListUsers() []response.UserResponse {
	panic("unimplemented")
}

// UpdateUser implements UserService
func (*UserServiceImpl) UpdateUser(request request.UserRequest) {
	panic("unimplemented")
}
