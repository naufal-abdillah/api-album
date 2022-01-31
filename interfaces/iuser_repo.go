package interfaces

import "example/web-service-gin/models"

type IUserRepo interface {
	RepoRegister(user models.User) (int64, error)
	RepoLogin(user models.User) (int64, error)
	UserExists(email string) (bool, error)
}
