package application

import (
	"github.com/joho/godotenv"
	"os"
	"rent/src/datasource"
)

var (
	DbConf datasource.DbConfig
	Port   string
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Can't load .env file")
	}

	DbConf = datasource.NewDbConfig(os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DP_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"),
	)

	Port = os.Getenv("APP_PORT")
}
