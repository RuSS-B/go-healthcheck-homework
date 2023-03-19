package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8000

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router(),
	}

	log.Printf("Starting server on port: %d", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		response := struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		}

		data, _ := json.Marshal(response)

		_, _ = w.Write(data)
	})

	return mux
}
