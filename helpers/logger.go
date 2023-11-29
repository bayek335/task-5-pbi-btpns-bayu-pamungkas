package helpers

import (
	"io"
	"os"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/sirupsen/logrus"
)

func setLogging() *logrus.Logger {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
		},
	}
	log.SetFormatter(jsonFormatter)
	log.SetLevel(logrus.DebugLevel)
	log.Out = os.Stdout
	return log
}

func UserLogging(userLog *app.UserLogging) {
	log := setLogging()
	file, err := os.OpenFile("user.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	if err == nil {
		log.Out = io.MultiWriter(file, os.Stdout)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	if userLog.Level == "info" {
		log.WithFields(logrus.Fields{
			"after_update":  userLog.AfterUpdate,
			"before_update": userLog.BeforeUpdate,
			"action":        userLog.Action,
			"user":          userLog.User,
			"user_id":       userLog.UserID,
		}).Info(userLog.Message)
	} else if userLog.Level == "error" {
		log.WithFields(logrus.Fields{
			"after_update":  userLog.AfterUpdate,
			"before_update": userLog.BeforeUpdate,
			"action":        userLog.Action,
			"user":          userLog.User,
			"user_id":       userLog.UserID,
		}).Error(userLog.Message)
	}
}

func ErrorLogging(userID, action, msg string, err error) {
	log := setLogging()
	file, errFile := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	if errFile == nil {
		log.Out = io.MultiWriter(file, os.Stdout)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.WithFields(logrus.Fields{
		"action":  action,
		"error":   err.Error(),
		"user_id": userID,
	}).Error(msg)
}
