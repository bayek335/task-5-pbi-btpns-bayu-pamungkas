package services

import (
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/models"
	"github.com/google/uuid"
)

type PhotosService interface {
	GetPhotos(user_ID uuid.UUID) ([]app.PhotosJson, error)
	CreatePhoto(photoRequest *app.PhotosRequestCreate) (*app.Photos, error)
	UpdatePhoto(photoRequest *app.PhotosRequestUpdate, photo_ID, user_ID uuid.UUID) (*app.Photos, error)
	DeletePhoto(photo_ID, user_ID uuid.UUID) (*app.Photos, error)
	FindPhotoByID(photo_ID uuid.UUID) (*app.Photos, error)
}

type photoService struct {
	model models.PhotosModel
}

func NewPhoto(model models.PhotosModel) *photoService {
	return &photoService{model}
}

const photosTable string = "photos"

func (service *photoService) FindPhotoByID(photo_ID uuid.UUID) (*app.Photos, error) {

	photo, err := service.model.FindPhotoByID(photo_ID)
	if err != nil {
		err := helpers.ErrorDatabase(err, photosTable)
		return nil, err
	}
	return photo, nil
}

func (service *photoService) GetPhotos(user_ID uuid.UUID) ([]app.PhotosJson, error) {
	photos, err := service.model.GetPhotos(user_ID)
	if err != nil {
		err := helpers.ErrorDatabase(err, photosTable)
		return nil, err
	}
	return photos, nil
}

func (service *photoService) CreatePhoto(photoRequest *app.PhotosRequestCreate) (*app.Photos, error) {
	user_ID, _ := uuid.Parse(photoRequest.UserId)
	photo := &app.Photos{
		ID:       photoRequest.ID,
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		ImageUrl: photoRequest.ImageUrl,
		UserId:   user_ID,
	}
	photo, err := service.model.CreatePhoto(photo)
	if err != nil {
		err := helpers.ErrorDatabase(err, photosTable)
		return nil, err
	}
	return photo, nil
}

func (service *photoService) UpdatePhoto(photoRequest *app.PhotosRequestUpdate, photo_ID, user_ID uuid.UUID) (*app.Photos, error) {
	photo := &app.Photos{
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		ImageUrl: photoRequest.ImageUrl,
	}
	photo, err := service.model.UpdatePhoto(photo, photo_ID, user_ID)
	if err != nil {
		err := helpers.ErrorDatabase(err, photosTable)
		return nil, err
	}
	return photo, nil
}

func (service *photoService) DeletePhoto(photo_ID, user_ID uuid.UUID) (*app.Photos, error) {

	photo, err := service.model.DeletePhoto(photo_ID, user_ID)
	if err != nil {
		err := helpers.ErrorDatabase(err, photosTable)
		return nil, err
	}
	return photo, nil
}
