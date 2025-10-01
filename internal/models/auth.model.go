package models

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Response
	Token string `json:"token" example:"example : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
