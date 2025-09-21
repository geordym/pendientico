package main

import (
	"log"

	application_factory "github.com/geordym/pendientico/application/factory"
	environment_configuration "github.com/geordym/pendientico/infraestructure/configuration/environment"

	adapters "github.com/geordym/pendientico/infraestructure/adapters/authentication/keycloack"
	"github.com/geordym/pendientico/infraestructure/adapters/persistence/postgres/configuration"
	postgres_factory "github.com/geordym/pendientico/infraestructure/adapters/persistence/postgres/factory"
	"github.com/geordym/pendientico/infraestructure/database/setup"
	"github.com/geordym/pendientico/infraestructure/http/handler"
	routes "github.com/geordym/pendientico/infraestructure/http/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	environment := environment_configuration.LoadEnvironment()

	gormFactory, err := configuration.NewGormFactory(environment)
	if err != nil {
		log.Fatal("Ocurrio un error al inicializar gorm")
	}

	setup.InitDB(*environment)

	authenticationProviderCommunication, err := adapters.NewKeycloakAdapterFromEnv(*environment)
	if err != nil {
		log.Fatal("fallo inicializar el keycloack, el cliente")
	}

	userRepository := postgres_factory.NewUserRepository(gormFactory.DB)
	createUserUseCase := application_factory.NewCreateUserUseCase(userRepository, authenticationProviderCommunication)

	//security.InitKeycloak()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	userHandler := handler.NewUserHandler(*createUserUseCase)

	routes.Init(e, userHandler)
	e.Logger.Fatal(e.Start(":8085"))
}
