package helpers

import (
	"math/rand"
	"strings"
)

func GenerateRandomString(num int) string {
	letter := "ABCDRFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvxyz1234567890!@#$%^&,."
	letters := strings.Split(letter, "")
	var ranStr []string

	for i := 0; i < num; i++ {
		ranStr = append(ranStr, letters[rand.Intn(len(letters)-1)])
	}

	return strings.Join(ranStr, "")
}
