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

func TestUpdateUser(t *testing.T) {
	router := setRouter()
	db := DatabaseConnection()
	router.PUT("/v1/users/:id", controllers.NewUser(services.NewUser(models.NewUser(db))).UpdateUser)
	ID := uuid.New()

	t.Run("update user success", func(t *testing.T) {
		var username, email, password string = "Bayu Pamungkas update", "bayu@gmail.com", "111111"

		user := &app.UsersUpdateRequest{
			Username: username,
			Email:    email,
			Password: password,
		}
		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest("PUT", "/v1/users/"+ID.String(), bytes.NewBuffer(jsonValue))

		userResponse := &app.UsersResponse{
			Success: true,
			Message: "success",
		}
		userResponse.Data.ID = ID
		userResponse.Data.Username = username
		userResponse.Data.Email = email
		userResponse.Data.IsActive = false

		json, _ := json.Marshal(userResponse)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(json), w.Body.String())
	})

	t.Run("update user fail", func(t *testing.T) {
		// var ID, username, email, password string = "a", "Bayu Pamungkas", "bayugmail.com", "123123"
		userTest := []map[string]string{
			{
				"ID":       ID.String(),
				"username": "Bayu Pamungkas",
				"email":    "111111",
				"password": "111111",
				"eMessage": "field email must be type of email",
			},
			{
				"ID":       ID.String(),
				"username": "",
				"email":    "111111",
				"password": "111111",
				"eMessage": "field username is required",
			},
			{
				"ID":       ID.String(),
				"username": "Bayu Pamungkas",
				"email":    "bayu@gmail.com",
				"password": "11111",
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
			req, _ := http.NewRequest("PUT", "/v1/users/"+ID.String(), bytes.NewBuffer(jsonValue))

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
