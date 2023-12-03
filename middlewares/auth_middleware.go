package middlewares

import (
	"errors"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/gin-gonic/gin"
)

func AuthJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user, err := helpers.GetCookie(c)

		// // check cookie user if nil ask to login
		// if err != nil {
		// 	err := errors.New("unauthorize, login first")
		// 	response, httpCode := helpers.ErrorResponse(err)
		// 	c.JSON(httpCode, response)
		// 	c.Abort()
		// 	return
		// }

		// get auth token from header section
		tokenString := c.Request.Header.Get("Authorization")

		// check token with helpers
		if !helpers.ValidateJwtToken(tokenString, c.Param("id")) {
			err := errors.New("unauthorize, login first")
			response, httpCode := helpers.ErrorResponse(err)
			c.JSON(httpCode, response)
			c.Abort()
			return
		}

		c.Next()

	}
}
