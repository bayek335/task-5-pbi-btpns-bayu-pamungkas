package helpers

import (
	"encoding/json"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Cookie is not used, i could not delete cause it can be a cheat sheet
func SetCookieUser(c *gin.Context, user *app.Users) {
	json, _ := json.Marshal(user)
	c.SetCookie("x-user", string(json), 3600*6, "/", "localhost", true, false)
}

func GetCookie(c *gin.Context) (*app.Users, error) {
	userCookie, err := c.Cookie("x-user")
	if err != nil {
		ErrorLogging("", "get cookie", "cookie nil", err)
		return nil, err
	}

	var User *app.Users

	if err := json.Unmarshal([]byte(userCookie), &User); err != nil {
		ErrorLogging("", "parse json", "authjwttoken middleware", err)
		return nil, err
	}

	return User, nil
}

func SetSession(c *gin.Context, user *app.Users) {

	session := sessions.Default(c)
	userJson, _ := json.Marshal(user)
	session.Set("user", string(userJson))
	err := session.Save()
	if err != nil {
		ErrorLogging("", "SetSession", "error while save session", err)
	}
}
