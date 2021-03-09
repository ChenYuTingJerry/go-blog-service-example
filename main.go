package main

import (
	"github/go-programming-tour-book/blog-service/internal/routes"
	"net/http"
	"time"
)

func main() {
	router := routes.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
