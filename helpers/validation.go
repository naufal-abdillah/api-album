package helpers

import (
	"errors"
	"example/web-service-gin/models"
	"regexp"
)

func ValidateUser(input models.User) error {
	if input.Email == "" {
		return errors.New("email required")
	}
	if len(input.Password) <= 0 {
		return errors.New("password required")
	}
	if input.Name == "" {
		return errors.New("name required")
	}
	regex, err := regexp.Compile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	if err != nil {
		return errors.New("error checking email validity")
	}
	var isMatch bool = regex.MatchString(input.Email)
	if !isMatch {
		return errors.New("use a valid email")
	}
	return err
}
