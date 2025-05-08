package controllers

import (
	"net/http"

	"ReturnsPersonApi/app/dataBaseConfig"
	"ReturnsPersonApi/app/models"

	"github.com/gin-gonic/gin"
)

func GetPersons(c *gin.Context) {
	rows, err := dataBaseConfig.DB.Query("SELECT id, nome, email FROM persons")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var persons []models.Person

	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.Id, &person.Nome, &person.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		persons = append(persons, person)
	}

	c.JSON(http.StatusOK, persons)
}

func CreatePerson(c *gin.Context) {
	var newPerson models.Person

	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json inválido"})
		return
	}

	query := "INSERT INTO persons(nome, email) VALUES ($1, $2) RETURNING id"
	err := dataBaseConfig.DB.QueryRow(query, newPerson.Nome, newPerson.Email).Scan(&newPerson.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPerson)
}
