package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		log.Fatal(err)
	}
	pass = string(hashPass)
	return pass
}

func ComparePassword(hashPass string, pass string) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass)); err != nil {
		return false
	}
	return true
}
