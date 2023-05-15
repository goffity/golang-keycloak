package main

import (
	"identity-keycloak/database"
	v1 "identity-keycloak/member/v1"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database, err := database.InitMySQLDatabase()
	if err != nil {
		panic("Cannot connect database")
	}

	database.AutoMigrate(&v1.Member{})

	repository := v1.NewRepository(database)
	service := v1.NewService(repository)
	api := v1.NewRegisterController(service)

	e.GET("/:id", api.GetEventByID)

	e.Logger.Fatal(e.Start(":1323"))
}
