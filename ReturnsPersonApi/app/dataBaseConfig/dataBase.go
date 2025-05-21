package dataBaseConfig

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conect() {
	dataCon := "host=localhost user=postgres password=datapostgres port=5432 dbname=dados_pessoas sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dataCon), &gorm.Config{})
	if err != nil {
		log.Fatal("erro ao conectar no banco de dados", err)
	}
	fmt.Println("Conexão com o banco estabelicida com sucesso!")
}

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
