package controllers

import (
	"net/http"
	"strconv"

	"github.com/CarlosBarbosaGomes/structure-app-mvc/config"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/helpers"
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

// @BasePath /api/v1

// @Tags		Users
// Create User	godoc
// @Summary 		Create users
// @Description 	Resource to save user and your credentials
// @Param			user body request.UserRequest{} true "Create User"
// @Produce			application/json
// @User			Users
// @Success			201 {object} response.UserResponse{}
// @Router			/users [post]
func (controller *UserController) Create(ctx *gin.Context) {
	request := request.UserRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		loggerMethods().Errorf("Error to parse request %v", request)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request, erro to parse request body"})
	}
	controller.Service.CreateUser(request)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created with success"})
}

// @Tags			Users
// List all users	godoc
// @Summary 		List all users
// @Description 	Resource to list all users
// @Produce			application/json
// @User			Users
// @Success			200 {object} response.UserResponse{}
// @Router			/users [get]
func (controller *UserController) List(ctx *gin.Context) {
	response := controller.Service.ListUsers()

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

// @Tags			Users
// Get user by id	godoc
// @Summary 		Return user by id
// @Param			id path string true "get user"
// @Description 	Resource to return user by id
// @Produce			application/json
// @User			Users
// @Success			200 {object} response.UserResponse{}
// @Router			/users/{id} [get]
func (controller *UserController) GetUser(ctx *gin.Context) {
	idUser := ctx.Param("id")
	id, err := strconv.Atoi(idUser)
	helpers.ErrorPanic(err)

	response := controller.Service.GetUserById(uint(id))
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

// @Tags			Users
// Remove user  	godoc
// @Param			id path string true "delete user"
// @Summary 		Delete user by id
// @Description 	Resource to remove user by id
// @Produce			application/json
// @User			Users
// @Success			204
// @Router			/users/{id} [delete]
func (controller *UserController) DeleteUser(ctx *gin.Context) {
	idUser := ctx.Param("id")
	id, err := strconv.Atoi(idUser)
	helpers.ErrorPanic(err)

	controller.Service.DeleteUser(uint(id))

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusNoContent, gin.H{"message": "User deleted with success"})
}

// @Tags			Users
// Update User	   	godoc
// @Summary 		Update user
// @Param			id path string true "update user by id"
// @Description 	Resource to update infos users
// @Param			user body request.UserRequestUpdate{} true "Update User"
// @Produce			application/json
// @User			Users
// @Success			200
// @Router			/users [put]
func (controller *UserController) UpdateUser(ctx *gin.Context) {
	requestBody := request.UserRequestUpdate{}
	err := ctx.ShouldBindJSON(&requestBody)
	helpers.ErrorPanic(err)

	idUser := ctx.Param("id")
	id, err := strconv.Atoi(idUser)
	helpers.ErrorPanic(err)
	requestBody.ID = uint(id)

	controller.Service.UpdateUser(requestBody)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated with success"})
}
