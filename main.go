package main

import (
	"Hamar/apis"
	"Hamar/database"

	"github.com/labstack/echo/v4"
)

var e = &apis.E{
	Cho: echo.New(),
	DB:  database.CreateDBConnectionPool(),
}

func init() {
	err := database.CheckProvisioning().DBProvisioningPipeline()
	if err != nil {
		panic(err)
	}
}

func main() {
	e.Cho.Logger.Info("Startup application finished")
	apis.RegisterRoutes(e)
	e.Cho.Start(":8080")
}
