package service

import (
	"time"

	"github.com/CarlosBarbosaGomes/structure-app-mvc/config"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers/request"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers/response"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/helpers"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/models"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/repositories"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic(err.Error())
	}

	userSave := models.Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hash),
	}

	service.IUserRepository.CreateUser(userSave)
	loggerMethods().Info("Service:: User save with success")
}

// DeleteUser implements UserService
func (service *UserServiceImpl) DeleteUser(id uint) {
	service.IUserRepository.DeleteUser(id)
}

// GetUserById implements UserService
func (service *UserServiceImpl) GetUserById(id uint) response.UserResponse {
	result := service.IUserRepository.GetUserById(id)
	if result.ID == 0 {
		loggerMethods().Infof("User with id %v not found", id)
	}
	response := response.UserResponse{
		ID:       result.ID,
		Name:     result.Name,
		Email:    result.Email,
		CreateAt: result.CreatedAt,
		UpdateAt: result.UpdatedAt,
	}

	return response
}

// GetUserByEmail implements UserService
func (service *UserServiceImpl) GetUserByEmail(email string) response.UserResponseLogin {
	result := service.IUserRepository.GetUserByEmail(email)
	if result.ID == 0 {
		loggerMethods().Infof("User with e-mail %v not found", email)
	}
	response := response.UserResponseLogin{
		ID:       result.ID,
		Name:     result.Name,
		Email:    result.Email,
		Password: result.Password,
	}

	return response
}

// ListUsers implements UserService
func (service *UserServiceImpl) ListUsers() []response.UserResponse {
	results := service.IUserRepository.ListUsers()

	var users []response.UserResponse

	for _, value := range results {
		user := response.UserResponse{
			ID:       value.ID,
			Name:     value.Name,
			Email:    value.Email,
			CreateAt: value.CreatedAt,
			UpdateAt: value.CreatedAt,
		}
		users = append(users, user)
	}
	return users
}

// UpdateUser implements UserService
func (service *UserServiceImpl) UpdateUser(user request.UserRequestUpdate) {
	result := service.IUserRepository.GetUserById(user.ID)

	result.Name = user.Name
	result.Email = user.Email
	result.Password = user.Password
	result.UpdatedAt = time.Now()

	service.IUserRepository.UpdateUser(result)
}
