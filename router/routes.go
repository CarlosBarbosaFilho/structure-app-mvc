package router

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/docs"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/database"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/middlewares"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/repositories"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func configrureRoutersController() *gin.Engine {

	// Configure database and validator
	databse := database.DB
	validate := validator.New()

	//Configure router default
	router := gin.Default()
	baseRouter := router.Group("/api/v1")

	//Configure swagger
	configureSwagger(router)

	//Configure routes users
	configureRoutesUser(baseRouter, databse, validate)

	//Configure routes login
	configureRoutesLogin(baseRouter, databse, validate)

	//Validate login user
	configureRoutesValidate(baseRouter, databse, validate)

	return router

}

func configureSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json"))
}

func configureRoutesLogin(baseRouter *gin.RouterGroup, databse *gorm.DB, validate *validator.Validate) {
	usersLogin := baseRouter.Group("/login")

	repository := repositories.NewUserRepositoryImpl(databse)
	service := service.NewUserServiceImpl(repository, validate)
	controllerLogin := controllers.NewUserLoginController(service)

	usersLogin.POST("", controllerLogin.LoginUser)
}

func configureRoutesUser(baseRouter *gin.RouterGroup, databse *gorm.DB, validate *validator.Validate) {
	usersRouter := baseRouter.Group("/users")

	repository := repositories.NewUserRepositoryImpl(databse)
	service := service.NewUserServiceImpl(repository, validate)
	controller := controllers.NewUserController(service)

	usersRouter.POST("", controller.Create)
	usersRouter.GET("", controller.List)
	usersRouter.GET("/:id", controller.GetUser)
	usersRouter.DELETE("/:id", controller.DeleteUser)
	usersRouter.PUT("", controller.UpdateUser)

}

func configureRoutesValidate(baseRouter *gin.RouterGroup, databse *gorm.DB, validate *validator.Validate) {
	validateLogin := baseRouter.Group("/validate")

	repository := repositories.NewUserRepositoryImpl(databse)
	service := service.NewUserServiceImpl(repository, validate)
	controller := controllers.NewUserController(service)

	validateLogin.GET("", middlewares.RequireAtuthentication, controller.Validate)
}
