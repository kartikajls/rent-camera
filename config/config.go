package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Loading .env")
	}
}

func ConnectDatabase() (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetBrevoAPIKey() string {
	return os.Getenv("BREVO_API_KEY")
}

func GetEmailFrom() string {
	return os.Getenv("EMAIL_FROM")
}
