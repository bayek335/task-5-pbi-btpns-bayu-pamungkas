package middlewares

import (
	"errors"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/gin-gonic/gin"
)

func AuthJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := "test"

		tokenString := c.Request.Header.Get("Authorization")

		if !helpers.ValidateJwtToken(tokenString, userID) {
			err := errors.New("unauthorize, login first")
			response, httpCode := helpers.ErrorResponse(err)
			c.JSON(httpCode, response)
			c.Abort()
			return
		}

		c.Next()

	}
}
