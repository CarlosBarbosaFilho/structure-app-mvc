package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func configureRoutes(router *gin.Engine) {
	//logger := config.NewLogger("Logging - Routes Configuration")
	v1 := router.Group("/api/v1")
	{
		v1.GET("/home", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
		})
		// v1.GET("/projects", handler.Projects)
		// v1.GET("/projects/:id", handler.GetProject)
		// v1.POST("/projects", handler.CreateProject)
		// v1.PUT("/projects", handler.UpdateProject)
		// v1.DELETE("/projects/:id", handler.DeleteProject)
	}

}
