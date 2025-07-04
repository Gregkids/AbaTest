package repository

import (
	"database/sql"
	"errors"

	"aba.technical.test/models"
	_ "github.com/lib/pq"
)

type UserSQL struct {
	DB *sql.DB
}

func (q *UserSQL) Login(req models.UserAuth) (string, error) {
	var uid string
	var pw string

	// Login Query
	Q := `
		SELECT
			u.id_user,
			u.password_user
		FROM users u WHERE username=?;
	`
	err := q.DB.QueryRow(Q, req.Email).Scan(&uid, &pw)

	switch {
	case err == sql.ErrNoRows:
		return "", errors.New("user not found")
	case err != nil:
		return "", err
	}
	// if !shared.MatchPassword(pw, req.Pass) {
	// 	return "", errors.New("password not match")
	// }

	// token, err := shared.CreateToken(uid)
	// if err != nil {
	// 	return "", errors.New("failed creating token")
	// }

	return "token", nil
}
