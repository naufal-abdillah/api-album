package handler

import (
	"example/web-service-gin/interfaces"
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"
	"example/web-service-gin/services"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HandlerRegisterUser(c *gin.Context) {
	var input models.User
	// var input map[string]string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var IUserRepo interfaces.IUserRepo = repositories.UserRepo{}
	userExists, err := (IUserRepo.UserExists(input.Email))
	// userExists, err := (IUserRepo.UserExists(input["email"]))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Email already registered",
		})
		return
	} else {
		var IUserService interfaces.IUserService = services.UserService{}
		// var Service services.UserService
		IUserService.ServicesRegister(input)
		token, err := createToken(input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error creating token",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "Success",
			"token":  token,
		})
		return
	}

	// fmt.Print(user["email"])
}
func createToken(input models.User) (string, error) {
	var SecretKey string = "RaHaSia"
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    input.Email,
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	return token, err

}
