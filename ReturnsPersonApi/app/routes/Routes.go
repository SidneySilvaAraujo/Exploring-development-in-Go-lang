package routes

import (
	"ReturnsPersonApi/app/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	engine := gin.Default()
	log.Println("setup routes")

	engine.GET("/allPersons", controllers.GetPersons)
	engine.POST("/createPerson", controllers.CreatePerson)

	return engine
}
