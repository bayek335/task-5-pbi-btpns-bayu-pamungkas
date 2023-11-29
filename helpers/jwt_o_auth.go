package helpers

import (
	"errors"
	"time"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(user *app.Users) string {
	mySigningKey := []byte(app.GetConfig().JwtKey)

	expTime := time.Now().Add(6 * time.Hour)
	// Create the Claims
	claims := &app.JwtClaims{
		UserID:   user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		panic("failed signed token to string")
	}

	return tokenString

}

func ValidateJwtToken(tokenString string, userID string) bool {

	token, err := jwt.ParseWithClaims(tokenString, &app.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.GetConfig().JwtKey), nil
	})
	action := "ValidateJwtToken"
	if err != nil {
		msg := "JwtToken does not match"
		ErrorLogging(userID, action, msg, err)
		return false
	}
	if !token.Valid {
		msg := "JwtToken expired"
		err := errors.New("JwtToken already expired")
		ErrorLogging(userID, action, msg, err)
		return false
	}
	return true
}
