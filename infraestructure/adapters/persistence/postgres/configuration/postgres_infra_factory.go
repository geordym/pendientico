package configuration

import (
	"fmt"

	configuration "github.com/geordym/pendientico/infraestructure/configuration/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormFactory struct {
	DB *gorm.DB
}

func NewGormFactory(env *configuration.Environment) (*GormFactory, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC search_path=%s",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
		env.DBSSLMode,
		env.DBSchema,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &GormFactory{DB: db}, nil
}
