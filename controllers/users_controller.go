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

func (control *userController) UpdateUser(c *gin.Context) {
	var userRequest app.UsersUpdateRequest
	user_ID, _ := uuid.Parse(c.Param("id"))

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		err := helpers.ErrorValidation(err)
		response, _ := helpers.ErrorResponse(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user, err := control.service.UpdateUser(userRequest, user_ID)
	if err != nil {
		response, httpCode := helpers.ErrorResponse(err)
		c.JSON(httpCode, response)
		return
	}
	response, _ := helpers.SuccessResponse(user)
	c.JSON(http.StatusOK, response)

}
