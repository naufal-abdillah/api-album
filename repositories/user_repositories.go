package repositories

import (
	"example/web-service-gin/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
}

func (R *UserRepo) RepoRegister(user models.User) (int64, error) {
	var id int64
	query := `INSERT INTO tb_user (name, email, password) VALUES (?, ?, ?)`
	row, err := db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	id, err = row.LastInsertId()
	return id, err
}

func (R *UserRepo) RepoLogin(input models.User) (int64, error) {
	user, err := QueryUser(input.Email)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	id, _ := strconv.Atoi(user.ID)
	return int64(id), err
}

func (R *UserRepo) UserExists(email string) (bool, error) {
	userExists := db.QueryRow(`SELECT EXISTS (SELECT * FROM tb_user WHERE email =? )`, email)
	var result bool
	err := userExists.Scan(&result)
	return result, err
}

func QueryUser(email string) (user models.User, err error) {
	rows, err := db.Query("SELECT * FROM tb_user WHERE email=?", email)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return user, err
		}
	}
	return user, err
}
