package handler

import (
	"example/web-service-gin/helpers"
	"example/web-service-gin/interfaces"
	"example/web-service-gin/models"
	"example/web-service-gin/services"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HandlerRegisterUser(c *gin.Context) {
	var input models.User
	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = helpers.ValidateRegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}
	var IUserService interfaces.IUserService = services.UserService{}

	id, err := IUserService.ServicesRegister(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}
	token, err := createToken(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"id":     id,
		"token":  token,
	})
}
func HandlerLoginUser(c *gin.Context) {
	var input models.User
	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
