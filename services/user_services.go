package services

import (
	"errors"
	"example/web-service-gin/interfaces"
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func (S UserService) ServicesRegister(input models.User) (int64, error) {
	var IUserRepo interfaces.IUserRepo = &repositories.UserRepo{}
	userExists, err := (IUserRepo.UserExists(input.Email))
	if err != nil {
		return 0, err
	}
	if userExists {
		return 0, errors.New("email already used")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(password),
	}
	var Repo repositories.UserRepo
	id, err := Repo.RepoRegister(user)
	if err != nil {
		return 0, errors.New("database error")
	}
	return id, nil
}

func (S UserService) ServicesLogin(input models.User) (int64, error) {
	var IUserRepo interfaces.IUserRepo = &repositories.UserRepo{}
	userExists, err := (IUserRepo.UserExists(input.Email))
	if err != nil {
		return 0, err
	}
	if !userExists {
		return 0, errors.New("account is not registered")
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	return IUserRepo.RepoLogin(user)

}
