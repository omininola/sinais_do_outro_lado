package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/a-h/templ"
	views "sinais/views"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.Handle("GET /", templ.Handler(views.HomePage()))

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
