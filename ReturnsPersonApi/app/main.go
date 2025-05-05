package main

import (
	"ReturnsPersonApi/app/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	engine := gin.Default()
	log.Printf("setup routes")

	engine.GET("/person", controllers.GetPerson)

	return engine
}

func main() {
	r := setupRoutes()
	r.Run()
}
