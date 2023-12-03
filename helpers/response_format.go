package helpers

import (
	"strings"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
)

func errorCode(err error) int {

	var httpCode int
	if strings.Contains(err.Error(), "already") {
		httpCode = 409
	} else if strings.Contains(err.Error(), "not") {
		httpCode = 404
	} else if strings.Contains(err.Error(), "unauthorize") {
		httpCode = 401
	} else if strings.Contains(err.Error(), "internal") {
		httpCode = 500
	} else {
		httpCode = 400
	}

	return httpCode
}

func ErrorResponse(err error) (*app.ErrorResponse, int) {
	response := &app.ErrorResponse{
		Success: false,
		Message: err.Error(),
	}
	httpCode := errorCode(err)
	return response, httpCode
}

func SuccessResponse(user *app.Users, action string) (*app.UsersResponse, int) {

	response := &app.UsersResponse{
		Success: true,
		Message: "user successfully " + action,
	}
	response.Data.ID = user.ID
	response.Data.Username = user.Username
	response.Data.Email = user.Email
	response.Data.IsActive = user.IsActive
	return response, 200
}

func LoginSuccessResponse(user *app.Users, token string) *app.UsersLoginResponse {
	response := &app.UsersLoginResponse{}
	response.Success = true
	response.Message = "user successfully loged in"
	response.Data.User.ID = user.ID
	response.Data.User.Username = user.Username
	response.Data.User.Email = user.Email
	response.Data.User.IsActive = user.IsActive
	response.Data.User.ActivatedAt = user.ActivatedAt
	response.Data.User.CreatedAt = user.CreatedAt
	response.Data.User.UpdatedAt = user.UpdatedAt
	response.Data.Token = token

	return response
}

func UsersProfileResponse(user *app.Users, action string) (*app.UsersProfileResponse, int) {

	response := &app.UsersProfileResponse{
		Success: true,
		Message: "user successfully " + action,
	}
	response.Data.ID = user.ID
	response.Data.Username = user.Username
	response.Data.Email = user.Email
	response.Data.IsActive = user.IsActive
	response.Data.ActivatedAt = user.ActivatedAt
	response.Data.CreatedAt = user.CreatedAt
	response.Data.UpdatedAt = user.UpdatedAt
	return response, 200
}

/*
| Photo response
*/

func AllPhotoSuccessResponse(photo []app.PhotosJson, action string) (*app.AllPhotosResponse, int) {
	response := &app.AllPhotosResponse{
		Success: true,
		Message: "photos successfully " + action,
	}
	response.Data = photo
	return response, 201
}

func PhotoSuccessResponse(photo *app.Photos, action string) (*app.PhotosResponse, int) {
	response := &app.PhotosResponse{
		Success: true,
		Message: "photo successfully " + action,
	}
	response.Data.ID = photo.ID.String()
	response.Data.Title = photo.Title
	response.Data.Caption = photo.Caption
	response.Data.ImageUrl = photo.ImageUrl
	response.Data.ProfileImage = photo.ProfileImage
	response.Data.UserId = photo.UserId.String()
	response.Data.CreatedAt = photo.CreatedAt
	return response, 201
}
