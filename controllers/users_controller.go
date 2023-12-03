package controllers

import (
	"net/http"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userController struct {
	service services.UsersService
}

func NewUser(service services.UsersService) *userController {
	return &userController{service}
}

func userLogInfo(user *app.Users, action, mes string) {
	// initiate for user logging
	userLog := &app.UserLogging{
		Level:   "info",
		UserID:  user.ID.String(),
		User:    user.Username,
		Action:  action,
		Message: "user successfully " + mes,
	}
	// logging user
	helpers.UserLogging(userLog)
}

/*
/	update user by user id
*/
func (control *userController) FindUserByID(c *gin.Context) {

	var user_ID uuid.UUID
	userSession := parsingUserSession(c)
	if userSession != nil {
		user_ID = userSession.ID
	} else {
		user_ID, _ = uuid.Parse(c.Param("id"))
	}

	user, err := control.service.FindUserByID(user_ID)
	if err != nil {
		response, httpCode := helpers.ErrorResponse(err)
		c.JSON(httpCode, response)
		return
	}
	// logger
	userLogInfo(user, "FindUserByID", "taken")

	response, _ := helpers.UsersProfileResponse(user, "taken")
	c.JSON(http.StatusOK, response)
}

/*
/	update user by user id
*/
func (control *userController) UpdateUser(c *gin.Context) {
	var userRequest app.UsersUpdateRequest
	user_ID, _ := uuid.Parse(c.Param("id"))

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		err := helpers.ErrorValidation(err)
		response, _ := helpers.ErrorResponse(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// hash password using bcrypt
	userRequest.Password = helpers.HashPassword(userRequest.Password)

	user, err := control.service.UpdateUser(userRequest, user_ID)
	if err != nil {
		response, httpCode := helpers.ErrorResponse(err)
		c.JSON(httpCode, response)
		return
	}

	// logger
	userLogInfo(user, "UpdateUser", "updated")

	response, _ := helpers.SuccessResponse(user, "updated")
	c.JSON(http.StatusOK, response)

}

/*
/	delete user by user id
*/
func (control *userController) DeleteUser(c *gin.Context) {
	user_ID, _ := uuid.Parse(c.Param("id"))
	user, err := control.service.DeleteUser(user_ID)
	if err != nil {
		response, httpCode := helpers.ErrorResponse(err)
		c.JSON(httpCode, response)
		return
	}

	// logger
	userLogInfo(user, "DeleteUser", "deleted")

	response, _ := helpers.SuccessResponse(user, "deleted")
	c.JSON(http.StatusOK, response)

}
