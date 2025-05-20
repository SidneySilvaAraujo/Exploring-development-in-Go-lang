package controllers

import (
	"net/http"
	"strconv"

	"ReturnsPersonApi/app/dataBaseConfig"
	"ReturnsPersonApi/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPersons(c *gin.Context) {
	var persons []models.Person
	if err := dataBaseConfig.DB.Find(&persons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, persons)
}

func CreatePerson(c *gin.Context) {
	var newPerson models.Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json inválido"})
		return
	}

	if err := dataBaseConfig.DB.Create(&newPerson).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, newPerson)
}

func UpdatePerson(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido!"})
		return
	}
	var existingPerson models.Person
	if err := dataBaseConfig.DB.First(&existingPerson, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "desculpe, pessoa não encontrada."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao procurar pessoas."})
		}
		return
	}

	var UpdatedPerson models.Person
	if err := c.ShouldBindJSON(&UpdatedPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jjson inválido."})
		return
	}

	existingPerson.Nome = UpdatedPerson.Nome
	existingPerson.Email = UpdatedPerson.Email
	if err := dataBaseConfig.DB.Save(&existingPerson).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar a pessoa."})
		return
	}

	c.JSON(http.StatusOK, existingPerson)
}
