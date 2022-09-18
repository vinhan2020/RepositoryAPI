package config

import (
	"github.com/subosito/gotenv"
	"os"
)

var (
	Host     string
	User     string
	Password string
	Port     string
	DBName   string
)

func Init() {
	_ = gotenv.Load()
	Host = os.Getenv("HOST")
	User = os.Getenv("USER_NAME")
	Password = os.Getenv("PASSWORD")
	Port = os.Getenv("PORT")
	DBName = os.Getenv("DB_NAME")
}
