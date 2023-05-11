package service

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers/request"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers/response"
)

type UserService interface {
	CreateUser(request request.UserRequest)
	ListUsers() []response.UserResponse
	DeleteUser(id uint)
	GetUserById(id uint) response.UserResponse
	UpdateUser(request request.UserRequestUpdate)
	GetUserByEmail(email string) response.UserResponseLogin
}
