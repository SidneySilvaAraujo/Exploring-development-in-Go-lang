package main

import (
	"apiv02/app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/person", handlers.ReturnsPersonHandler)
	r.Run()
}
