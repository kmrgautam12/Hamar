package main

import (
	"Hamar/apis"
	"Hamar/database"

	"github.com/labstack/echo/v4"
)

var e = echo.New()

func init() {
	err := database.CheckProvisioning().DBProvisioningPipeline()
	if err != nil {
		panic(err)
	}
}

func main() {
	e.Logger.Info("Startup application finished")
	apis.RegisterRoutes(e)
	e.Start(":8080")
}
