package services

import (
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func (S UserService) ServicesRegister(input models.User) {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	// password, _ := bcrypt.GenerateFromPassword(user["password"], 12)
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
	}
	// fmt.Print(user.ID)
	var Repo repositories.UserRepo
	Repo.RepoRegister(user)
}
