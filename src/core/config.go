package core

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	JWTSecret          string
	BcryptCost         int
	FCMProjectID       string
	FCMCredentialsFile string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("  No se encontró .env, usando variables de entorno del sistema")
	}

	bcryptCost, err := strconv.Atoi(getEnv("BCRYPT_COST", "14"))
	if err != nil {
		log.Fatal("BCRYPT_COST debe ser un número entero")
	}

	return &Config{
		DBHost:             getEnv("DB_HOST", ""),
		DBPort:             getEnv("DB_PORT", "3306"),
		DBUser:             getEnv("DB_USER", ""),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", ""),
		JWTSecret:          getEnv("JWT_SECRET", ""),
		BcryptCost:         bcryptCost,
		FCMProjectID:       getEnv("FCM_PROJECT_ID", ""),
		FCMCredentialsFile: getEnv("FCM_CREDENTIALS_FILE", ""),
	}
}

// Retorna el valor de la variable o un default si no existe
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	if defaultValue == "" {
		log.Fatalf(" Variable de entorno requerida no encontrada: %s", key)
	}
	return defaultValue
}