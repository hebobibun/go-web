package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About page"))
	}

	// function outside of main
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/selena", selenaHandler)

	// function inside of main
	mux.HandleFunc("/about", aboutHandler)

	// function inside of HandleFunc
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Contact page"))
	})

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hellow home, I'm learning Golang web"))
}

func selenaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selena from Land of Dawn"))
}