package app

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username    string    `db:"username"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	IsActive    bool      `db:"is_active"`
	ActivatedAt time.Time `db:"activated_at"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type UsersCreateRequest struct {
	ID       uuid.UUID
	Username string `json:"username" binding:"required,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"min=6"`
}

type UsersUpdateRequest struct {
	Username string `json:"username" binding:"required,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"min=6"`
}

type UsersLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UsersResponse struct {
	Success bool   `json:"succes"`
	Message string `json:"message"`
	Data    struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
		IsActive bool      `json:"is_active"`
	} `json:"data"`
}

type UsersLoginResponse struct {
	Success bool   `json:"succes"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}
