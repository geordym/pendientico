package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	environment_configuration "github.com/geordym/pendientico/infraestructure/configuration/environment"
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

func GetDB(env *environment_configuration.Environment) *sql.DB {
	once.Do(func() {
		// Validar que todos los campos requeridos estén presentes
		if env.DBHost == "" || env.DBPort == "" || env.DBUser == "" ||
			env.DBPassword == "" || env.DBName == "" || env.DBSSLMode == "" || env.DBSchema == "" {
			log.Fatal("Faltan variables de configuración de la base de datos en Environment")
		}

		// Construir la conexión
		connStr := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			env.DBHost, env.DBPort, env.DBUser, env.DBPassword, env.DBName, env.DBSSLMode, env.DBSchema,
		)

		var err error
		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal("Error conectando a la base de datos:", err)
		}

		// Configuración de pool
		DB.SetMaxOpenConns(25)
		DB.SetMaxIdleConns(5)

		// Probar la conexión
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
