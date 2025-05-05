package controllers

import (
	"database/sql"
	"net/http"

	"ReturnsPersonApi/app/dataBaseConfig"
	"ReturnsPersonApi/app/models"

	"github.com/gin-gonic/gin"
)

func GetPerson(c *gin.Context) {
	var person models.Person
	err := dataBaseConfig.DB.QueryRow("SELECT id, nome, email FROM persons LIMIT 1").Scan(&person.Id, &person.Nome, &person.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "nem uma pessoa encontrada"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, person)
}
