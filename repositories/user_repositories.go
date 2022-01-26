package repositories

import (
	"example/web-service-gin/models"
	"net/http"
)

type UserRepo struct {
}

func (R UserRepo) RepoRegister(user models.User) int {
	entry := `INSERT INTO tb_user (id, name, email, password) VALUES (?, ?, ?, ?)`
	db.MustExec(entry, user.ID, user.Name, user.Email, user.Password)

	return http.StatusOK
}

func (R UserRepo) UserExists(email string) (bool, error) {
	userExists := db.QueryRow(`SELECT EXISTS (SELECT * FROM tb_user WHERE email =? )`, email)
	var result bool
	err := userExists.Scan(&result)
	return result, err
}
