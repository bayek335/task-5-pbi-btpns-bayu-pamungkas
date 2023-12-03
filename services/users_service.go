package services

import (
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/models"
	"github.com/google/uuid"
)

type UsersService interface {
	FindUserByEmail(email string) (*app.Users, error)
	CreateUser(userRequest app.UsersCreateRequest) (*app.Users, error)
	UpdateUser(userRequest app.UsersUpdateRequest, user_ID uuid.UUID) (*app.Users, error)
	DeleteUser(user_ID uuid.UUID) (*app.Users, error)
	FindUserByID(user_ID uuid.UUID) (*app.Users, error)
}

type userService struct {
	model models.UsersModel
}

func NewUser(model models.UsersModel) *userService {
	return &userService{model}
}

const userTable string = "users"

func (service *userService) FindUserByEmail(email string) (*app.Users, error) {
	user, err := service.model.FindUserByEmail(email)
	// check error from sql
	if err != nil {
		err = helpers.ErrorLogin()
		return nil, err
	}
	return user, nil
}

func (service *userService) CreateUser(userRequest app.UsersCreateRequest) (*app.Users, error) {
	// parsing format user request struct to users struct
	user := &app.Users{
		ID:       userRequest.ID,
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	user, err := service.model.CreateUser(user)
	// check error from sql
	if err != nil {
		err = helpers.ErrorDatabase(err, userTable)
		return nil, err
	}
	return user, nil
}

func (service *userService) UpdateUser(userRequest app.UsersUpdateRequest, user_ID uuid.UUID) (*app.Users, error) {
	User := &app.Users{}

	// parsing format user request struct to users struct
	User.Username = userRequest.Username
	User.Email = userRequest.Email
	User.Password = userRequest.Password

	user, err := service.model.UpdateUser(User, user_ID)
	// check error from sql
	if err != nil {
		err = helpers.ErrorDatabase(err, userTable)
		return nil, err
	}
	return user, nil
}

func (service *userService) DeleteUser(user_ID uuid.UUID) (*app.Users, error) {
	user, err := service.model.DeleteUser(user_ID)
	// check error from sql
	if err != nil {
		err = helpers.ErrorDatabase(err, userTable)
		return nil, err
	}
	return user, nil
}

func (service *userService) FindUserByID(user_ID uuid.UUID) (*app.Users, error) {
	user, err := service.model.FindUserByID(user_ID)
	// check error from sql
	if err != nil {
		err = helpers.ErrorDatabase(err, userTable)
		return nil, err
	}
	return user, nil
}
