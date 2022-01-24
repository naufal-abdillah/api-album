package handler

import (
	"example/web-service-gin/repositories"
	"example/web-service-gin/services"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HandlerRegisterUser(c *gin.Context) {
	var input map[string]string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	userExists, err := (repositories.UserExists(input["email"]))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if userExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Email already registered",
		})
	} else {
		var Service services.UserService
		Service.ServicesRegister(input)
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
	}

	// fmt.Print(user["email"])
}
func createToken(input map[string]string) (string, error) {
	var SecretKey string = "RaHaSia"
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    input["Email"],
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	return token, err

}
