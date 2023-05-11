package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CarlosBarbosaGomes/structure-app-mvc/controllers/response"
	"github.com/CarlosBarbosaGomes/structure-app-mvc/initializer/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAtuthentication(ctx *gin.Context) {
	fmt.Println("In Middleware")
	//Get cookie of request
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return ([]byte(os.Getenv("SECRET_ACCESS"))), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		//Find user with user sub
		var user response.UserResponseLogin
		//initializers.DB.First(&user, claims["sub"])
		database.DB.Raw("SELECT * FROM users WHERE id = ?", claims["sub"]).Scan(&user)
		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		//Attach to request on context
		ctx.Set("user", user)

		//Continue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
