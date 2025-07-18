package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	SSLMode           string
	TimeZone          string
	JWTSecret         string
	JWTExpired        int
	JWTRefreshExpired int
	HTTPPort          int
}

func LoadEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file not found, using system environment variables")
	}

	return &Env{
		DBHost:            os.Getenv("DB_HOST"),
		DBPort:            os.Getenv("DB_PORT"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		SSLMode:           os.Getenv("DB_SSLMODE"),
		TimeZone:          os.Getenv("DB_TIMEZONE"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTExpired:        getInt("JWT_EXPIRED", 3600),
		JWTRefreshExpired: getInt("JWT_REFRESH_EXPIRED", 7200),
		HTTPPort:          getInt("HTTP_PORT", 8080),
	}
}

func getInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	parsed, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return parsed
}
