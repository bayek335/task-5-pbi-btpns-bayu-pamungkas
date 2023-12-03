package models

import (
	"errors"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UsersModel interface {
	FindUserByEmail(email string) (*app.Users, error)
	CreateUser(user *app.Users) (*app.Users, error)
	UpdateUser(user *app.Users, user_ID uuid.UUID) (*app.Users, error)
	DeleteUser(user_ID uuid.UUID) (*app.Users, error)
	FindUserByID(user_ID uuid.UUID) (*app.Users, error)
}

type userModel struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *userModel {
	return &userModel{db}
}

func (model *userModel) FindUserByEmail(email string) (*app.Users, error) {
	var User *app.Users

	// take error if exist from database while find user by email
	err := model.db.Where("email=?", email).First(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (model *userModel) CreateUser(user *app.Users) (*app.Users, error) {

	// take error if exist from database while create user account
	err := model.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (model *userModel) UpdateUser(user *app.Users, user_ID uuid.UUID) (*app.Users, error) {
	var User app.Users
	result := model.db.Model(&User).Clauses(clause.Returning{}).Where("id = ?", user_ID).Updates(user)
	if result.RowsAffected < 1 {
		err := errors.New(gorm.ErrRecordNotFound.Error())
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &User, nil
}

func (model *userModel) DeleteUser(user_ID uuid.UUID) (*app.Users, error) {
	var User app.Users
	result := model.db.Clauses(clause.Returning{}).Where("id = ?", user_ID).Delete(&User)
	if result.RowsAffected < 1 {
		err := errors.New(gorm.ErrRecordNotFound.Error())
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &User, nil
}

func (model *userModel) FindUserByID(user_ID uuid.UUID) (*app.Users, error) {
	var User app.Users
	result := model.db.Find(&User, user_ID)
	if result.RowsAffected < 1 {
		err := errors.New(gorm.ErrRecordNotFound.Error())
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &User, nil
}
