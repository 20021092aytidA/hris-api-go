package env

import (
	"os"

	"github.com/joho/godotenv"
)

type structENV struct {
	DBName string
	DBHost string
	DBPort string
	DBUser string
	DBPass string

	APIPort string
}

var ENV structENV

func LoadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	ENV.DBName = os.Getenv("DB_NAME")
	ENV.DBHost = os.Getenv("DB_HOST")
	ENV.DBPort = os.Getenv("DB_PORT")
	ENV.DBUser = os.Getenv("DB_USER")
	ENV.DBPass = os.Getenv("DB_PASS")
	ENV.APIPort = os.Getenv("API_PORT")

	return nil
}
