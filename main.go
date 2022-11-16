package main

import (
	"go-web/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About page"))
	}

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/selena", handler.SelenaHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Contact page"))
	})

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

