package main

import (
	"github.com/gorilla/mux"

	"lobster/internal/handler"
	"lobster/internal/repository"
	"lobster/internal/service"

	"fmt"
	"net/http"
	"time"
)

func main() {
	storage := repository.NewLinkStore()
	s := &service.LinkService{Storage: storage}
	h := &handler.LinkHandler{Service: s}

	r := mux.NewRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	r.HandleFunc("/api/links", h.CreateLink).Methods("POST")
	r.HandleFunc("/api/links/{id}", h.GetLink).Methods("GET")

	fmt.Println("http://localhost:8080")

	server.ListenAndServe()
}
