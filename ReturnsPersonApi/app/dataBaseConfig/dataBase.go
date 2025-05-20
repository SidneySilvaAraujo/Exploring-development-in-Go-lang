package dataBaseConfig

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conect() {
	dataCon := "user=postgres password=datapostgres dbname=dados_pessoas sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dataCon), &gorm.Config{})
	if err != nil {
		log.Fatal("erro ao conectar no banco de dados", err)
	}
	fmt.Println("Conexão com o banco estabelicida com sucesso!")
}
