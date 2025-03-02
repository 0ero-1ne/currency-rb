package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		log.Printf("Environment variable %s is empty. Using default value: %s", key, defaultValue)
		return defaultValue
	}

	return value
}
