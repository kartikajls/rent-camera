package dto

import "time"

type UserResponse struct {
	UserID        int64     `json:"user_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Role          string    `json:"role"`
	DepositAmount float64   `json:"deposit_amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type RegisterResponse struct {
	UserID        int64     `json:"user_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	DepositAmount float64   `json:"deposit_amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type LoginResponse struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	JWTToken string `json:"jwt_token,omitempty"`
}
