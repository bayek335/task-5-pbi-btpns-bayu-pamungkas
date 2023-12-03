package app

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username    string    `db:"username"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	IsActive    bool      `db:"is_active"`
	ActivatedAt time.Time `db:"activated_at"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type UsersProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID          uuid.UUID `json:"id"`
		Username    string    `json:"username"`
		Email       string    `json:"email"`
		IsActive    bool      `json:"is_active"`
		ActivatedAt time.Time `json:"activated_at"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"data"`
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
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
		IsActive bool      `json:"is_active"`
	} `json:"data"`
}

type UsersLoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		User struct {
			ID          uuid.UUID `json:"id"`
			Username    string    `json:"username"`
			Email       string    `json:"email"`
			IsActive    bool      `json:"is_active"`
			ActivatedAt time.Time `json:"activated_at"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
		} `json:"user"`
		Token string `json:"token"`
	} `json:"data"`
}
