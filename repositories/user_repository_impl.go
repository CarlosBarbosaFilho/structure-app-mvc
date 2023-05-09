package repositories

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/config"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) IUserRepository {
	return &UserRepositoryImpl{DB: db}
}

func loggerMethods() *config.Logger {
	logger := config.NewLogger("Logging - Routes Configuration")
	return logger
}

// CreateUser implements IUserRepository
func (repo *UserRepositoryImpl) CreateUser(user models.UserModel) {
	result := repo.DB.Create(&user)
	if result != nil {
		loggerMethods().Error("Error to create a new user")
	}
}

// DeleteUser implements IUserRepository
func (repo *UserRepositoryImpl) DeleteUser(id uint) {
	var user models.UserModel
	result := repo.DB.Delete(&user, id)
	if result != nil {
		loggerMethods().Errorf("Error to delete user with id %v", id)
	}

}

// GetUserById implements IUserRepository
func (repo *UserRepositoryImpl) GetUserById(id uint) models.UserModel {
	var user models.UserModel
	result := repo.DB.Find(&user, id)
	if result != nil {
		loggerMethods().Errorf("Error to find user with id %v", id)
	}
	return user
}

// ListUsers implements IUserRepository
func (repo *UserRepositoryImpl) ListUsers() []models.UserModel {
	var users []models.UserModel
	result := repo.DB.Select([]models.UserModel{}).Find(&users)
	if result != nil {
		loggerMethods().Error("Error to list all users")
	}
	return users
}

// UpdateUser implements IUserRepository
func (repo *UserRepositoryImpl) UpdateUser(user models.UserModel) {
	var userEdit = models.UserModel{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	result := repo.DB.Model(&user).Updates(userEdit)
	if result != nil {
		loggerMethods().Error("Error to find user")
	}

}
