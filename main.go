package main

import (
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
}
