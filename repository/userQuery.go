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

func (q *UserSQL) Login(req *models.UserCred) (string, error) {
	var uid, pw, role string

	// Login Query
	query := `
		SELECT
			u.user_id,
			u.password,
			u.role
		FROM public.users u WHERE email=$1;
	`
	err := q.DB.QueryRow(query, req.Email).Scan(&uid, &pw, &role)

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

func (q *UserSQL) GetUser(reqID string) ([]models.UserProfile, error) {
	ret := []models.UserProfile{}

	// Query Get User Profile
	query := `
	SELECT
		u.user_id,
		u.name,
		u.email,
		u.role
	FROM public.users u WHERE user_id=$1;
	`
	data := models.UserProfile{}
	err := q.DB.QueryRow(query, reqID).Scan(&data.UserID, &data.Username, &data.Email, &data.Role)

	if err == sql.ErrNoRows {
		return nil, errors.New("data not found")
	} else if err != nil {
		return nil, err
	}

	ret = append(ret, data)

	return ret, nil
}
