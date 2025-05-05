package dataBaseConfig

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Conect() {
	var err error
	connStr := "user=postgres dbname=dados_pessoas sslmode=disable password=datapostgres"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Conexão com o banco estabelecida com sucesso!")
}
