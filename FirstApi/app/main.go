package main

import (
	"log"
	"net/http"

	"FirstAPI/app/controllers"
)

func main() {
	http.HandleFunc("/hello", controllers.HelloHandler)
	log.Println("Servidor iniciado na porta: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
