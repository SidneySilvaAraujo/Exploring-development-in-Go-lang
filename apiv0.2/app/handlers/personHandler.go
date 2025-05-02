package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Nome  string `json:"nome"`
	Idade string `json:"idade"`
}

func ReturnsPersonHandler(c *gin.Context) {
	person := Person{
		Nome:  "Sidney",
		Idade: "20",
	}
	c.JSON(http.StatusOK, person)
}
