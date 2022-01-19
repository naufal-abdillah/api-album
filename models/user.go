package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"title"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}
