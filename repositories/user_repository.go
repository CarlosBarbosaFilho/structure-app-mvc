package repositories

import "github.com/CarlosBarbosaGomes/structure-app-mvc/models"

type IUserRepository interface {
	CreateUser(user models.Users)
	ListUsers() []models.Users
	DeleteUser(id uint)
	GetUserById(id uint) models.Users
	UpdateUser(user models.Users)
}
