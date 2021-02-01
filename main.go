package main

import (
	"go-dependency-injection/homePage"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "log ", log.LstdFlags|log.Ltime|log.Lshortfile)

	h := homePage.NewHandlers(logger)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		logger.Printf("Inside handler")
		rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello world"))
	})

	mux.Handle("/home", h)
	http.ListenAndServe(":8080", mux)
}
