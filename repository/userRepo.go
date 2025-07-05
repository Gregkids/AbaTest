package repository

import (
	"database/sql"
	"errors"

	"aba.technical.test/models"
	"aba.technical.test/service"
	_ "github.com/lib/pq"
)

type UserSQL struct {
	DB *sql.DB
}

func (q *UserSQL) Login(req *models.UserAuth) (string, error) {
	var uid, pw, role string

	// Login Query
	Q := `
		SELECT
			u.user_id,
			u.password,
			u.id_role
		FROM users u WHERE email=?;
	`
	err := q.DB.QueryRow(Q, req.Email).Scan(&uid, &pw, &role)

	switch {
	case err == sql.ErrNoRows:
		return "", errors.New("user not found")
	case err != nil:
		return "", err
	}

	if !service.MatchPassword(pw, req.Pass) {
		return "", errors.New("password not match")
	}

	token, err := service.GenerateToken(uid, role)
	if err != nil {
		return "", errors.New("failed creating token")
	}

	return token, nil
}
