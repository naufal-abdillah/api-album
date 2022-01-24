package services

import (
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func (S UserService) ServicesRegister(input map[string]string) {
	password, _ := bcrypt.GenerateFromPassword([]byte(input["password"]), 12)
	// password, _ := bcrypt.GenerateFromPassword(user["password"], 12)
	user := models.User{
		Name:     input["name"],
		Email:    input["email"],
		Password: password,
	}
	// fmt.Print(user.ID)
	var Repo repositories.UserRepo
	Repo.RepoRegister(user)
}
