package interfaces

import "example/web-service-gin/models"

type IUserService interface {
	ServicesRegister(input models.User)
}
