package main

import (
	"fmt"
	"log"
	"net/http"

	views "sinais/views"

	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.Handle("GET /", templ.Handler(views.HomePage()))
	mux.Handle("GET /alphabet", templ.Handler(views.AlphabetPage()))

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
