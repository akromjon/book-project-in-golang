package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string, d string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	v := os.Getenv(key)

	if v != "" {
		return v
	}

	return d
}

type App struct {
	Name  string
	Env   string
	Debug string
	Url   string
	Port  string
}

func NewApp() App {
	return App{
		Name:  Getenv("APP_NAME", "GOLANG"),
		Env:   Getenv("APP_ENV", "Local"),
		Debug: Getenv("APP_DEBUG", "false"),
		Url:   Getenv("APP_URL", "http://localhost"),
		Port:  Getenv("APP_PORT", "8080"),
	}
}
