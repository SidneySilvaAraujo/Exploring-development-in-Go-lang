package main

import (
	"ReturnsPersonApi/app/dataBaseConfig"
	"ReturnsPersonApi/app/routes"
)

func main() {
	r := routes.SetupRoutes()

	dataBaseConfig.Conect()

	r.Run()
}
