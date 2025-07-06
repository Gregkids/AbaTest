package models

type UserCred struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

type UserProfile struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
