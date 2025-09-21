package setup

import (
	"log"

	"github.com/geordym/pendientico/infraestructure/configuration/database"
	environment_configuration "github.com/geordym/pendientico/infraestructure/configuration/environment"
	"github.com/pressly/goose/v3"
)

func InitDB(environment environment_configuration.Environment) {
	db := database.GetDB(&environment)

	// Verificar conexión
	if err := db.Ping(); err != nil {
		log.Fatal("Error en la conexión:", err)
	}

	log.Println("Base de datos inicializada correctamente")

	// Ejecutar migraciones
	if err := goose.Up(db, "./infraestructure/database/migrations"); err != nil {
		log.Fatal("Error ejecutando migraciones:", err)
	}

	log.Println("Migraciones aplicadas correctamente ✅")
}
