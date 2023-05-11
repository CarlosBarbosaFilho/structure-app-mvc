package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := initializerControllers()
	router.Run()
}

func initializerControllers() *gin.Engine {
	router := configrureRoutersController()
	return router
}
