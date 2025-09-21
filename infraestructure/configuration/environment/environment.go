package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBSchema   string

	// Keycloak
	KeycloakURL          string
	KeycloakRealm        string
	KeycloakClientID     string
	KeycloakClientSecret string
}

func LoadEnvironment() *Environment {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar .env, usando variables del entorno")
	}

	return &Environment{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		DBSchema:   os.Getenv("DB_SCHEMA"),

		KeycloakURL:          os.Getenv("KEYCLOAK_URL"),
		KeycloakRealm:        os.Getenv("KEYCLOAK_REALM"),
		KeycloakClientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		KeycloakClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
	}
}
