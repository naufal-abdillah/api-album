package repositories

import (
	"example/web-service-gin/models"
)

func (R Repo) RepoRegister(user models.User) {
	// var exists bool = false
	// userExists := db.QueryRow(`SELECT EXISTS (SELECT * FROM tb_user WHERE (name, email)=(?,?) )`, user.Name, user.Email)
	// err := (userExists.Scan(0))
	// if err != nil
	entry := `INSERT INTO tb_user (id, name, email, password) VALUES (?, ?, ?, ?)`
	db.MustExec(entry, user.ID, user.Name, user.Email, user.Password)
}
