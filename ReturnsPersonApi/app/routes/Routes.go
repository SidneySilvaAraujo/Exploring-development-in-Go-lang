package routes

import (
	"ReturnsPersonApi/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	engine.GET("/allPersons", controllers.GetPersons)
	engine.POST("/createPerson", controllers.CreatePerson)
	engine.PUT("/update/:id", controllers.UpdatePerson)
	engine.DELETE("/delete/:id", controllers.DeletePerson)
}
