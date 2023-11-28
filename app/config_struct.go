package app

import "os"

type Config struct {
	ServerHost string
	ServerPort string

	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

func GetConfig() *Config {
	return &Config{
		ServerHost: os.Getenv("SERVER_HOST"),
		ServerPort: os.Getenv("SERVER_PORT"),

		DatabaseHost:     os.Getenv("DB_HOST"),
		DatabasePort:     os.Getenv("DB_PORT"),
		DatabaseUser:     os.Getenv("DB_USER"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
		DatabaseName:     os.Getenv("DB_NAME"),
	}
}
