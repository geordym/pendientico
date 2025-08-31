package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB   *sql.DB
	once sync.Once
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró archivo .env, usando variables de entorno del sistema")
	}
}

func GetDB() *sql.DB {
	once.Do(func() {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		sslmode := os.Getenv("DB_SSLMODE")
		schema := os.Getenv("DB_SCHEMA")
		if schema == "" {
			schema = "public"
		}

		if host == "" {
			host = "localhost"
		}
		if port == "" {
			port = "5432"
		}
		if sslmode == "" {
			sslmode = "disable"
		}

		connStr := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			host, port, user, password, dbname, sslmode, schema,
		)

		var err error
		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal("Error conectando a la base de datos:", err)
		}

		DB.SetMaxOpenConns(25)
		DB.SetMaxIdleConns(5)

		if err = DB.Ping(); err != nil {
			log.Fatal("Error haciendo ping a la base de datos:", err)
		}

		log.Println("Conexión a la base de datos establecida exitosamente")
	})

	return DB
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
