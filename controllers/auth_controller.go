package controllers

import (
	"net/http"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type authController struct {
	service services.UsersService
}

func NewAuth(service services.UsersService) *authController {
	return &authController{service}
}

/*
/	create user account
*/
func (control *authController) Register(c *gin.Context) {
	var userRequest app.UsersCreateRequest

	// validating user request with user request struct
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		err := helpers.ErrorValidation(err)
		returnErr(c, err)
		return
	}

	// for unit testing, and comment if in deployment because unit testing also send a uuid
	if userRequest.ID.String() == "00000000-0000-0000-0000-000000000000" {
		userRequest.ID = uuid.New()
	}

	// deployment
	// userRequest.ID = uuid.New()

	// hash password using bcrypt
	userRequest.Password = helpers.HashPassword(userRequest.Password)
	// send user request to service
	user, err := control.service.CreateUser(userRequest)
	if err != nil {
		helpers.ErrorLogging("", "register", "register error", err)
		returnErr(c, err)
		return
	}
	response, _ := helpers.SuccessResponse(user, "created")
	c.JSON(http.StatusCreated, response)
}

/*
/	user login, generate token
*/
func (control *authController) Login(c *gin.Context) {
	var userRequest app.UsersLoginRequest

	// check validation error
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		err := helpers.ErrorValidation(err)
		returnErr(c, err)
		return
	}

	// check user is exist by email
	user, err := control.service.FindUserByEmail(userRequest.Email)
	if err != nil {
		helpers.ErrorLogging("", "login", "find user error", err)
		returnErr(c, err)
		return
	}

	// matching password between password user request and password in database
	if !helpers.ComparePassword(user.Password, userRequest.Password) {
		err := helpers.ErrorLogin()
		helpers.ErrorLogging(user.ID.String(), "login", "compare password error", err)
		returnErr(c, err)
		return
	}

	// initiate for user logging
	userLog := &app.UserLogging{
		Level:   "info",
		UserID:  user.ID.String(),
		User:    user.Username,
		Action:  "Login",
		Message: "user successfully loged in",
	}
	// generate jwt token
	jwtToken := helpers.GenerateJwtToken(user)

	// helpers.SetCookieUser(c, user)
	helpers.SetSession(c, user)

	// store user logging
	helpers.UserLogging(userLog)
	response := helpers.LoginSuccessResponse(user, jwtToken)
	c.JSON(http.StatusOK, response)

}
