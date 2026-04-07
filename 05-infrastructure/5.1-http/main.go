package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println("Can't send response: ", err)
			}
		}
	})

	mux.HandleFunc("GET /api/v1/users/{id}", UserID)

	mux.HandleFunc("POST /api/v1/tokens", TokenHand)

	log.Println("http server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
