package controllers

import (
	"net/http"

	"github.com/CarlosBarbosaGomes/structure-app-mvc/config"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/request"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{Service: service}
}

func loggerMethods() *config.Logger {
	logger := config.NewLogger("Logging - Routes Configuration")
	return logger
}

func (controller *UserController) Create(ctx *gin.Context) {
	request := request.UserRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		loggerMethods().Errorf("Error to parse request %v", request)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request, erro to parse request body"})
	}
	controller.Service.CreateUser(request)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{})
}
