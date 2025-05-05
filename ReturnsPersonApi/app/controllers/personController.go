package controllers

import (
	"net/http"

	"ReturnsPersonApi/app/models"

	"github.com/gin-gonic/gin"
)

func GetPerson(c *gin.Context) {
	person := models.Person{
		Id:    21000000,
		Nome:  "Satoshi",
		Email: "satoshi@nacamoto.com",
	}
	c.JSON(http.StatusOK, person)
}
