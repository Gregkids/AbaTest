package models

type UserAuth struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}
