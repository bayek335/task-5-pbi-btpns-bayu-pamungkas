package helpers

import (
	"strings"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
)

func errorCode(err error) int {

	var httpCode int
	if strings.Contains(err.Error(), "already exist") {
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

func LoginSuccessResponse(token string) *app.UsersLoginResponse {
	response := &app.UsersLoginResponse{}
	response.Success = true
	response.Message = "user successfully loged in"
	response.Data.Token = token

	return response
}
