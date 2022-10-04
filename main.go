package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	port := "8080"

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}
