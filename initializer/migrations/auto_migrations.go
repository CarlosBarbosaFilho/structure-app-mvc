package migrations

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/database"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/models"
)

func AutoMigrations() {
	database.DB.Table("users").AutoMigrate(&models.UserModel{})
}
