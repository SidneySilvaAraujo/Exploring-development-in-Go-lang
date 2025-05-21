package main

import (
	"ReturnsPersonApi/app/dataBaseConfig"
	"ReturnsPersonApi/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dataBaseConfig.Conect()

	r := gin.Default()

	r.Use(dataBaseConfig.DBMiddleware(dataBaseConfig.DB))

	routes.SetupRoutes(r)

	r.Run()
}
