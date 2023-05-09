package service

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/request"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/response"
)

type UserService interface {
	CreateUser(request request.UserRequest)
	ListUsers() []response.UserResponse
	DeleteUser(id uint)
	GetUserById(id uint) response.UserResponse
	UpdateUser(request request.UserRequest)
}
