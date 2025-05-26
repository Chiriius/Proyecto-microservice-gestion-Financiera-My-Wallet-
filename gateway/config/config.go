package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                  string
	AccountServiceAddress string
	UserServiceAddress    string
}

func Load() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}

	return &Config{
		Port:                  getEnv("PORT_GATEWAY", "8080"),
		AccountServiceAddress: getEnv("ACCOUNT_SERVICE_ADDRESS", "localhost:50052"),
		UserServiceAddress:    getEnv("USER_SERVICE_ADDRESS", "localhost:50051"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
