package app

import (
	"time"

	"github.com/google/uuid"
)

type Photos struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title        string    `db:"title"`
	Caption      string    `db:"caption"`
	ImageUrl     string    `db:"image_url"`
	UserId       uuid.UUID `db:"user_id" gorm:"foreignKey"`
	ProfileImage bool      `db:"profile_image"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type PhotosRequestCreate struct {
	ID       uuid.UUID
	Title    string
	Caption  string `form:"caption"`
	ImageUrl string
	UserId   string `form:"user_id" binding:"required"`
}

type PhotosRequestUpdate struct {
	Title    string
	Caption  string `form:"caption"`
	ImageUrl string
}

type PhotosJson struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Caption      string    `json:"caption"`
	ImageUrl     string    `json:"image_url"`
	UserId       string    `json:"user_id"`
	ProfileImage bool      `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PhotosResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    PhotosJson `json:"data"`
}

type AllPhotosResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    []PhotosJson `json:"data"`
}
