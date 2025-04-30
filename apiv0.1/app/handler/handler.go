package handler

import (
	"fmt"
	"net/http"
)

func OiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá, mundo!")
}
