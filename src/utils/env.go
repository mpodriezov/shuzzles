package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return err
}

// Access the environment variables
// s3Bucket := os.Getenv("S3_BUCKET")
// secretKey := os.Getenv
