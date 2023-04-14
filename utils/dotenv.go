package authUtils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Failed to load File!")
	}
	return os.Getenv(key)
}