package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/controllers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/models"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var username, email, password string = "Bayu Pamungkas", "bayu@gmail.com", "123123"

func TestRegisterUser(t *testing.T) {
	router := setRouter()
	db := DatabaseConnection()
	router.POST("/v1/users/register", controllers.NewAuth(services.NewUser(models.NewUser(db))).Register)
	ID, _ := uuid.Parse(User_ID)
	t.Run("create user success", func(t *testing.T) {

		user := &app.UsersCreateRequest{
			ID:       ID,
			Username: username,
			Email:    email,
			Password: password,
		}

		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/v1/users/register", bytes.NewBuffer(jsonValue))

		userResponse := &app.UsersResponse{
			Success: true,
			Message: "success",
		}
		userResponse.Data.ID = user.ID
		userResponse.Data.Username = user.Username
		userResponse.Data.Email = user.Email
		userResponse.Data.IsActive = false

		json, _ := json.Marshal(userResponse)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, string(json), w.Body.String())
	})

	t.Run("create user fail", func(t *testing.T) {
		// var ID, username, email, password string = "a", "Bayu Pamungkas", "bayugmail.com", "123123"
		userTest := []map[string]string{
			{
				"ID":       ID.String(),
				"username": "Bayu Pamungkas",
				"email":    "bayugmail.com",
				"password": "123123",
				"eMessage": "field email must be type of email",
			},
			{
				"ID":       ID.String(),
				"username": "",
				"email":    "bayu@gmail.com",
				"password": "123123",
				"eMessage": "field username is required",
			},
			{
				"ID":       ID.String(),
				"username": "Bayu Pamungkas",
				"email":    "bayu@gmail.com",
				"password": "12312",
				"eMessage": "field password is min 6 characters",
			},
		}
		for _, val := range userTest {
			ID, _ := uuid.Parse(val["ID"])
			user := &app.UsersCreateRequest{
				ID:       ID,
				Username: val["username"],
				Email:    val["email"],
				Password: val["password"],
			}
			jsonValue, _ := json.Marshal(user)
			req, _ := http.NewRequest("POST", "/v1/users/register", bytes.NewBuffer(jsonValue))

			userResponse := &app.ErrorResponse{
				Success: false,
				Message: val["eMessage"],
			}

			json, _ := json.Marshal(userResponse)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, string(json), w.Body.String())
		}
	})
}

func TestLogin(t *testing.T) {
	router := setRouter()
	db := DatabaseConnection()
	router.POST("/v1/users/login", controllers.NewAuth(services.NewUser(models.NewUser(db))).Login)
	t.Run("login success", func(t *testing.T) {
		var actualResult *app.UsersLoginResponse

		userReq := &app.UsersLoginRequest{
			Email:    email,
			Password: password,
		}

		jsonReq, _ := json.Marshal(userReq)

		req, _ := http.NewRequest("POST", "/v1/users/login", bytes.NewBuffer(jsonReq))
		userResponse := &app.UsersLoginResponse{
			Success: true,
			Message: "user successfully loged in",
		}

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		wString := w.Body.String()

		if err := json.Unmarshal([]byte(wString), &actualResult); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, userResponse.Success, actualResult.Success)
		assert.Equal(t, userResponse.Message, actualResult.Message)
	})
	t.Run("login fail", func(t *testing.T) {
		userTest := []map[string]string{
			{
				"email":    "bayugmail.com",
				"password": "123123",
				"eMessage": "field email must be type of email",
			},
			{
				"email":    "bayu3@gmail.com",
				"password": "123123",
				"eMessage": "wrong email or password",
			},
			{
				"email":    "bayu@gmail.com",
				"password": "1",
				"eMessage": "wrong email or password",
			},
		}
		for _, val := range userTest {

			var expectedResult *app.ErrorResponse

			user := &app.UsersLoginRequest{
				Email:    val["email"],
				Password: val["password"],
			}
			jsonValue, _ := json.Marshal(user)
			req, _ := http.NewRequest("POST", "/v1/users/login", bytes.NewBuffer(jsonValue))

			userResponse := &app.ErrorResponse{
				Success: false,
				Message: val["eMessage"],
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			json.Unmarshal([]byte(w.Body.String()), &expectedResult)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, userResponse.Success, expectedResult.Success)
			assert.Equal(t, userResponse.Message, expectedResult.Message)
		}
	})
}
