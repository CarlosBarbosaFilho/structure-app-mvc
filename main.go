package main

import (
	"fmt"

	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/database"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/environment"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/logger"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/migrations"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/router"
)

func init() {
	database.DBConnection()
	environment.LoadEnvironmentVariables()
	migrations.AutoMigrations()
	logger.InitializerLogger()
	router.Initialize()
}

func main() {
	fmt.Println("My structure to mvc project in golang")
}
