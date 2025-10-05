package main

import (
	"log"

	application_factory "github.com/geordym/pendientico/application/factory"
	environment_configuration "github.com/geordym/pendientico/infraestructure/configuration/environment"
	"github.com/geordym/pendientico/infraestructure/configuration/security"

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
	workspaceRepository := postgres_factory.NewWorkspaceRepository(gormFactory.DB)
	workspaceUserRepository := postgres_factory.NewWorkspaceUsersRepository(gormFactory.DB)
	workspaceContactRepository := postgres_factory.NewWorkspaceContactRepository(gormFactory.DB)

	createUserUseCase := application_factory.NewCreateUserUseCase(userRepository, authenticationProviderCommunication)
	createWorkspaceUseCase := application_factory.NewCreateWorkspaceUseCase(workspaceRepository, workspaceUserRepository, userRepository, authenticationProviderCommunication)
	createWorkspaceContactUseCase := application_factory.NewCreateWorkspaceContactUseCase(workspaceContactRepository)

	security.InitKeycloak()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	userHandler := handler.NewUserHandler(*createUserUseCase)
	workspacesHandler := handler.NewWorkspaceHandler(*createWorkspaceUseCase, *createWorkspaceContactUseCase)

	routes.Init(e, userHandler, workspacesHandler)
	e.Logger.Fatal(e.Start(":8085"))
}
