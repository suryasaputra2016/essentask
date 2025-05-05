package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	IsVerified   bool   `json:"is_verified"`
	Role         string `json:"role"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}
