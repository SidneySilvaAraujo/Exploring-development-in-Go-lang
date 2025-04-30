package main

import (
	"log"
	"net/http"

	"apiv0.1/app/handler"
)

func main() {
	http.HandleFunc("/oi", handler.OiHandler)
	log.Println("Servidor iniciado na porta: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
