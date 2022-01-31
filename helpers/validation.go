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
	regex, err := regexp.Compile(`[a-z]+`)
	var isMatch bool = regex.MatchString(input.Email)
	if isMatch {
		return nil
	}
	return err
}
