package app

import (
	"fmt"
	"net/http"

	"github.com/Zawar-Ahmed10p/go_postgres/services"
)

func Run() {
	serve := http.NewServeMux()
	serve.HandleFunc("/", services.AddUser)

	fmt.Println("listening on port 8000")
	http.ListenAndServe(":8000", serve)
}
