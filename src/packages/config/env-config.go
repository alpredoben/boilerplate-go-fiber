package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfiguration(fileEnv string) {
	err := godotenv.Load(fileEnv)

	if err != nil {
		log.Fatalf("System can't load environment file (%s). Error steatment %v", fileEnv, err)
	}

	log.Print("Environment file read successfully...")

	LoadApplicationConfig()
	LoadDatabaseConfig()
}
