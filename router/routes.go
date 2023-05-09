package router

import (
	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers"
	"github.com/gin-gonic/gin"
)

func routerUserController(controller *controllers.UserController) *gin.Engine {
	router, baseRouter := confugurationControllerDefault()
	personasRouter := baseRouter.Group("/users")
	personasRouter.POST("", controller.Create)

	return router

}

func confugurationControllerDefault() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	baseRouter := router.Group("/api/v1")
	return router, baseRouter
}
