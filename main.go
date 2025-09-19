package main

import (
	"github.com/geordym/pendientico/configuration/security"
	"github.com/geordym/pendientico/database/setup"
	"github.com/geordym/pendientico/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	setup.InitDB()
	security.InitKeycloak()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.Init(e)
	e.Logger.Fatal(e.Start(":8085"))
}
