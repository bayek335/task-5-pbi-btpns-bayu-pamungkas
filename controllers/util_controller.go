package controllers

import (
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
/	create response json with error
*/
func returnErr(c *gin.Context, err error) {
	response, _ := helpers.ErrorResponse(err)
	c.JSON(http.StatusBadRequest, response)
}

func parsingUserSession(c *gin.Context) *app.UserSession {

	var userSession *app.UserSession
	// parsing session to struct
	userSessionInterface := sessions.Default(c).Get("user")
	if userSessionInterface == nil {
		return nil
	}
	userSessionString := userSessionInterface.(string)

	err := json.Unmarshal([]byte(userSessionString), &userSession)
	if err != nil {
		// logging
		helpers.ErrorLogging("", "FindUserByID", "parsing session to struct", err)
		err := errors.New("internal server error")
		response, httpCode := helpers.ErrorResponse(err)
		c.JSON(httpCode, response)
		return nil
	}

	return userSession

}

// validate file function
func validateFile(c *gin.Context, file *multipart.FileHeader, err error, user_ID string) bool {

	if err != nil {
		err := errors.New("the field image is required")
		returnErr(c, err)
		return false
	}
	// check file ext
	if !helpers.CheckFileExt(file) {
		err := errors.New("image file must jpg, jpeg or png")
		returnErr(c, err)
		return false
	}
	// filesize lowerthan 2mb
	if file.Size > 2048000 {
		err := errors.New("maximum size of the image file is 2048Kb")
		returnErr(c, err)
		return false
	}

	return true

}
