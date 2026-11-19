package main

import (
	"fmt"
	"log"
	"net/http"

	"sinais/lib"
	"sinais/views"

	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	fsAssets := http.FileServer(http.Dir("./assets/"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fsAssets))

	mux.Handle("GET /", templ.Handler(views.HomePage()))
	mux.Handle("GET /alphabet", templ.Handler(views.AlphabetPage()))

	mux.HandleFunc("GET /getRandomPhrase", func(w http.ResponseWriter, r *http.Request) {
		templ.Raw(lib.GetRandomPhrase()).Render(r.Context(), w)
	})

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
