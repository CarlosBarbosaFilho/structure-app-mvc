package repositories

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/config"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/helpers"
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
func (repo *UserRepositoryImpl) CreateUser(user models.Users) {
	result := repo.DB.Create(&user)
	if result == nil {
		loggerMethods().Errorf("Error to create a new user %v", result.Error)
		helpers.ErrorPanic(result.Error)
		return
	}
}

// DeleteUser implements IUserRepository
func (repo *UserRepositoryImpl) DeleteUser(id uint) {
	var user models.Users
	result := repo.DB.Delete(&user, id)
	helpers.ErrorPanic(result.Error)

}

// GetUserById implements IUserRepository
func (repo *UserRepositoryImpl) GetUserById(id uint) models.Users {
	var user models.Users
	result := repo.DB.Find(&user, id)
	helpers.ErrorPanic(result.Error)
	return user
}

// GetUserByEmail implements IUserRepository
func (repo *UserRepositoryImpl) GetUserByEmail(email string) models.Users {
	var user models.Users
	result := repo.DB.Where("email = ?", email).First(&user)

	helpers.ErrorPanic(result.Error)
	return user
}

// ListUsers implements IUserRepository
func (repo *UserRepositoryImpl) ListUsers() []models.Users {
	var users []models.Users
	result := repo.DB.Find(&users)
	helpers.ErrorPanic(result.Error)
	return users
}

// UpdateUser implements IUserRepository
func (repo *UserRepositoryImpl) UpdateUser(user models.Users) {
	var userEdit = models.Users{
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
	helpers.ErrorPanic(result.Error)

}
