package models

import (
	"errors"
	"fmt"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotosModel interface {
	GetPhotos(user_ID uuid.UUID) ([]app.PhotosJson, error)
	CreatePhoto(photo *app.Photos) (*app.Photos, error)
	UpdatePhoto(photo *app.Photos, photo_ID, User_ID uuid.UUID) (*app.Photos, error)
	DeletePhoto(photo_ID, user_ID uuid.UUID) (*app.Photos, error)
	FindPhotoByID(photo_ID uuid.UUID) (*app.Photos, error)
}
type photoModel struct {
	db *gorm.DB
}

func NewPhoto(db *gorm.DB) *photoModel {
	return &photoModel{db}
}

func (model *photoModel) FindPhotoByID(photo_ID uuid.UUID) (*app.Photos, error) {
	var Photo *app.Photos
	err := model.db.First(&Photo, photo_ID).Error
	if err != nil {
		return nil, err
	}
	return Photo, nil
}

func (model *photoModel) GetPhotos(user_ID uuid.UUID) ([]app.PhotosJson, error) {
	var Photo *app.Photos
	var photos []app.PhotosJson

	// return an array of photo
	err := model.db.Model(&Photo).Where("user_id =?", user_ID).Find(&photos).Error
	if err != nil {
		return nil, err
	}
	fmt.Println()
	return photos, nil
}

func (model *photoModel) CreatePhoto(photo *app.Photos) (*app.Photos, error) {
	err := model.db.Create(photo).Error
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (model *photoModel) UpdatePhoto(photo *app.Photos, photo_ID, user_ID uuid.UUID) (*app.Photos, error) {
	var Photo *app.Photos
	result := model.db.Model(&Photo).Where("id = ? AND user_id=?", photo_ID, user_ID).Updates(photo)
	if result.RowsAffected < 1 {
		err := errors.New(gorm.ErrRecordNotFound.Error())
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	photo.ID = photo_ID
	return photo, nil
}

func (model *photoModel) DeletePhoto(photo_ID, user_ID uuid.UUID) (*app.Photos, error) {
	var Photo *app.Photos
	result := model.db.Model(&Photo).Clauses(clause.Returning{}).Where("user_id=?", user_ID).Delete(&Photo, photo_ID)
	if result.RowsAffected < 1 {
		err := errors.New(gorm.ErrRecordNotFound.Error())
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return Photo, nil
}
