package main

import (
	"fmt"
	"net/http"

	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/database"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/environment"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/logger"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/migrations"
	"github.com/gin-gonic/gin"
)

func init() {
	database.DBConnection()
	environment.LoadEnvironmentVariables()
	migrations.AutoMigrations()
	logger.InitializerLogger()

}

func main() {

	fmt.Println("My structure to mvc project in golang")
	router := gin.Default()
	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome MVC structure app"})
	})

	router.Run()
}
