package helpers

import (
	"mime/multipart"
	"path"
	"slices"
	"strconv"
	"time"
)

func GenerateFileImage(file *multipart.FileHeader) (fileName, fileUrl string) {

	ext := path.Ext(file.Filename)
	hours := time.Now().Hour()
	minutes := time.Now().Minute()
	seconds := time.Now().Second()

	randStr := GenerateRandomString(20)

	fileName = randStr + strconv.Itoa(hours) + strconv.Itoa(minutes) + strconv.Itoa(seconds) + ext
	fileUrl = "public/images/" + fileName
	return fileName, fileUrl
}

func CheckFileExt(file *multipart.FileHeader) bool {
	allowedFileExt := []string{".jpg", ".jpeg", ".png"}

	return slices.Contains(allowedFileExt, path.Ext(file.Filename))
}
