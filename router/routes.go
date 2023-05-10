package router

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/docs"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func routerUserController(controller *controllers.UserController) *gin.Engine {
	router, baseRouter := confugurationControllerDefault()

	usersRouter := baseRouter.Group("/users")
	usersRouter.POST("", controller.Create)
	usersRouter.GET("", controller.List)
	usersRouter.GET("/:id", controller.GetUser)
	usersRouter.DELETE("/:id", controller.DeleteUser)
	usersRouter.PUT("", controller.UpdateUser)

	return router

}

func confugurationControllerDefault() (*gin.Engine, *gin.RouterGroup) {

	//Configure router default
	router := gin.Default()
	baseRouter := router.Group("/api/v1")

	//Configure swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"))

	return router, baseRouter
}
