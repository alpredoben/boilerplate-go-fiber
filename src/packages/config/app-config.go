package configs

import (
	"os"
	"strconv"
	"time"
)

type AppEnv struct {
	Host        string
	Port        int
	Debug       bool
	ReadTimeout time.Duration

	JWTSecretKey    string
	JWTExpireMinute int
}

var app = &AppEnv{}

func GetApplicationConfig() *AppEnv {
	return app
}

func LoadApplicationConfig() {
	app.Host = os.Getenv("APP_HOST")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	app.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	timeOut, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	app.ReadTimeout = time.Duration(timeOut) * time.Second
	app.JWTSecretKey = os.Getenv("JWT_SECRET")
	app.JWTExpireMinute, _ = strconv.Atoi(os.Getenv("JWT_SECRET_EXPIRE_MINUTE_COUNT"))
}
