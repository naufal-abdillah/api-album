package repositories

import (
	"example/web-service-gin/models"
)

type UserRepo struct {
}

func (R UserRepo) RepoRegister(user models.User) (int64, error) {
	var id int64
	query := `INSERT INTO tb_user (name, email, password) VALUES (?, ?, ?)`
	row, err := db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	id, err = row.LastInsertId()
	return id, err
}

func (R UserRepo) UserExists(email string) (bool, error) {
	userExists := db.QueryRow(`SELECT EXISTS (SELECT * FROM tb_user WHERE email =? )`, email)
	var result bool
	err := userExists.Scan(&result)
	return result, err
}
