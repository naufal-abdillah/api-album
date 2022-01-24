package interfaces

import "example/web-service-gin/models"

type IUserRepo interface {
	RepoRegister(user models.User) int
	UserExists(email string) (bool, error)
}
