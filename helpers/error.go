package helpers

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorValidation(err error) error {

	for _, e := range err.(validator.ValidationErrors) {
		field := strings.ToLower(e.Field())
		newErr := "field " + field + " is " + e.ActualTag()
		if e.ActualTag() == "min" {
			err = errors.New(newErr + " 6 characters")
			return err
		} else if e.ActualTag() == "max" {
			err = errors.New(newErr + " 64 characters")
			return err
		} else if e.ActualTag() == "email" {
			err = errors.New("field " + field + " must be type of " + e.ActualTag())
			return err
		}
		err = errors.New(newErr)
		return err
	}
	return err
}

func ErrorDatabase(err error, tbl string) error {

	if tbl == "users" {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("email already exist")
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user does not exist")
		}
	}
	if tbl == "photos" {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("photo does not exist")
		}
	}

	return err
}

func ErrorLogin() error {
	return errors.New("wrong email or password")
}
