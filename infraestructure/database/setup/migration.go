package setup

import (
	"log"

	database "github.com/geordym/pendientico/configuration"
	"github.com/pressly/goose/v3"
)

func InitDB() {
	db := database.GetDB()

	// Verificar conexión
	if err := db.Ping(); err != nil {
		log.Fatal("Error en la conexión:", err)
	}

	log.Println("Base de datos inicializada correctamente")

	// Ejecutar migraciones
	if err := goose.Up(db, "./database/migrations"); err != nil {
		log.Fatal("Error ejecutando migraciones:", err)
	}

	log.Println("Migraciones aplicadas correctamente ✅")
}
