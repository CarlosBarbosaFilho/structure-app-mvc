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

// @title Demo API - Structure MVC
// @version 1.0
// @description Doc Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8888
// @BasePath /api/v1
// @schemes http
func main() {
	fmt.Println("My structure to mvc project in golang")
}
