package router

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/database"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/repositories"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Initialize() {
	router := initializerControllers()
	router.Run()
}

func initializerControllers() *gin.Engine {
	databse := database.DB
	validate := validator.New()

	//Configuration user controllers routes
	repository := repositories.NewUserRepositoryImpl(databse)
	service := service.NewUserServiceImpl(repository, validate)
	controller := controllers.NewUserController(service)
	router := routerUserController(controller)

	return router
}
