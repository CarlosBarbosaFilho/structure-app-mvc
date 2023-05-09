package repositories

import "github.com/CarlosBarbosaGomes/structure-app-mvc/models"

type IUserRepository interface {
	CreateUser(user models.UserModel)
	ListUsers() []models.UserModel
	DeleteUser(id uint)
	GetUserById(id uint) models.UserModel
	UpdateUser(user models.UserModel)
}
