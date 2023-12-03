package controllers

import (
	"net/http"
	"os"
	"strings"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type photosController struct {
	service services.PhotosService
}

func NewPhoto(service services.PhotosService) *photosController {
	return &photosController{service}
}

// get all photos by user id which user id get from session
func (control *photosController) GetPhotos(c *gin.Context) {
	user_ID := parsingUserSession(c).ID

	photos, err := control.service.GetPhotos(user_ID)
	if err != nil {
		returnErr(c, err)
		return
	}

	response, _ := helpers.AllPhotoSuccessResponse(photos, "taken")
	c.JSON(http.StatusCreated, response)

}

// create photo
func (control *photosController) CreatePhoto(c *gin.Context) {
	var photoRequest *app.PhotosRequestCreate

	// cause i try to access from front end so a make a decision
	var user_ID string
	userSession := parsingUserSession(c)
	if userSession != nil {
		user_ID = userSession.ID.String()
	} else {
		user_ID = c.Param("id")
	}

	// validating user request with user request struct
	if err := c.Bind(&photoRequest); err != nil {
		err := helpers.ErrorValidation(err)
		returnErr(c, err)
		return
	}

	// is file not nil
	file, err := c.FormFile("image")

	// validate image file
	if !validateFile(c, file, err, user_ID) {
		return
	}

	fileName, fileUrl := helpers.GenerateFileImage(file)

	photoRequest.ID = uuid.New()
	photoRequest.Title = fileName
	photoRequest.ImageUrl = "http://" + app.GetConfig().ServerHost + ":" + app.GetConfig().ServerPort + "/" + fileUrl

	// save file
	if err := c.SaveUploadedFile(file, fileUrl); err != nil {
		returnErr(c, err)
		return
	}
	// send user request to service
	photo, err := control.service.CreatePhoto(photoRequest)
	if err != nil {
		os.Remove(fileUrl)
		returnErr(c, err)
		return
	}

	response, _ := helpers.PhotoSuccessResponse(photo, "created")
	c.JSON(http.StatusCreated, response)
}

func (control *photosController) UpdatePhoto(c *gin.Context) {
	var photoRequest *app.PhotosRequestUpdate

	photo_ID, _ := uuid.Parse(c.Param("id"))

	// checking uer id where it is consume postman or FE
	var user_ID string
	var username string
	userSession := parsingUserSession(c)
	if userSession != nil {
		user_ID = userSession.ID.String()
		username = userSession.Username
	} else {
		user_ID = c.Param("id")
	}

	// validating user request with user request struct
	if err := c.Bind(&photoRequest); err != nil {
		err := helpers.ErrorValidation(err)
		returnErr(c, err)
		return
	}

	// is file not nil
	file, err := c.FormFile("image")

	// validate image file
	if !validateFile(c, file, err, user_ID) {
		return
	}

	fileName, fileUrl := helpers.GenerateFileImage(file)

	var cutSet string = "http://" + app.GetConfig().ServerHost + ":" + app.GetConfig().ServerPort + "/"
	photoRequest.Title = fileName
	photoRequest.ImageUrl = cutSet + fileUrl

	// get old file
	oldPhotos, err := control.service.FindPhotoByID(photo_ID)

	if err != nil {
		returnErr(c, err)
		return
	}

	// save file
	if err := c.SaveUploadedFile(file, fileUrl); err != nil {
		returnErr(c, err)
		return
	}
	// send user request to service
	photo, err := control.service.UpdatePhoto(photoRequest, photo_ID, userSession.ID)
	if err != nil {
		os.Remove(fileUrl)
		returnErr(c, err)
		return
	}

	// delete old image
	pathImage := strings.TrimLeft(oldPhotos.ImageUrl, cutSet)
	os.Remove("p" + pathImage)

	// initiate user log
	userLog := &app.UserLogging{
		Level:        "info",
		UserID:       user_ID,
		User:         username,
		Action:       "UpdatePhoto",
		BeforeUpdate: oldPhotos.Title,
		AfterUpdate:  photo.Title,
	}
	// logging
	helpers.UserLogging(userLog)

	photo.UserId = oldPhotos.UserId
	response, _ := helpers.PhotoSuccessResponse(photo, "update")
	c.JSON(http.StatusOK, response)
}

func (control *photosController) DeletePhoto(c *gin.Context) {
	photo_ID, _ := uuid.Parse(c.Param("id"))
	user_ID := parsingUserSession(c).ID
	// get old file
	oldPhotos, err := control.service.FindPhotoByID(photo_ID)
	if err != nil {
		returnErr(c, err)
		return
	}

	photo, err := control.service.DeletePhoto(photo_ID, user_ID)
	if err != nil {
		returnErr(c, err)
		return
	}

	// delete old image
	var cutSet string = "http://" + app.GetConfig().ServerHost + ":" + app.GetConfig().ServerPort + "/"
	pathImage := strings.TrimLeft(oldPhotos.ImageUrl, cutSet)
	os.Remove("p" + pathImage)

	response, _ := helpers.PhotoSuccessResponse(photo, "deleted")
	c.JSON(http.StatusOK, response)
}
