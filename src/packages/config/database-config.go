package configs

import (
	"os"
	"strconv"
	"time"
)

type DBEnv struct {
	Host                  string
	Port                  int
	SslMode               string
	Database              string
	Username              string
	Password              string
	Debug                 bool
	MaxOpenConnection     int
	MaxIdleConnection     int
	MaxLifeTimeConnection time.Duration
}

var db = &DBEnv{}

func GetDatabaseConfig() *DBEnv {
	return db
}

func LoadDatabaseConfig() {
	db.Host = os.Getenv("DB_HOSTNAME")
	db.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	db.Username = os.Getenv("DB_USERNAME")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Database = os.Getenv("DB_DATABASE")
	db.SslMode = os.Getenv("DB_SSL_MODE")
	db.Debug, _ = strconv.ParseBool(os.Getenv("DB_DEBUG"))
	db.MaxOpenConnection, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTION"))
	db.MaxIdleConnection, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTION"))
	lifeTime, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTION"))
	db.MaxLifeTimeConnection = time.Duration(lifeTime) * time.Second
}
