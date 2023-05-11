package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers/request"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserControllerLogin struct {
	Service service.UserService
}

func NewUserLoginController(service service.UserService) *UserControllerLogin {
	return &UserControllerLogin{Service: service}
}

// @BasePath /api/v1

// @Tags			Login
// Create User		godoc
// @Summary 		Login user
// @Description 	Resource to login user
// @Param			user body request.UserRequestLogin{} true "Login User"
// @Produce			application/json
// @User			User
// @Router			/login [post]
func (controller *UserControllerLogin) LoginUser(ctx *gin.Context) {

	// create a user request to parse values
	request := request.UserRequestLogin{}

	// validate the parse
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Erro to login access"})
		return
	}

	// search user on database
	var userLogin = controller.Service.GetUserByEmail(request.Email)

	// verifiy if user exists
	if userLogin.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Email or password invalid"})
		return
	}

	// compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(request.Password))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Error to compare passwords"})
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userLogin.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_ACCESS")))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "failed to create token"})
		return
	}
	//set token in cookie
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{})
}
