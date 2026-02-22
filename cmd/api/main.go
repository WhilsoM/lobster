package main

import (
	"fmt"
	"lobster/internal/handler"
	"lobster/internal/repository"
	"lobster/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	storage := repository.NewLinkStore()
	s := &service.LinkService{Storage: storage}
	h := &handler.LinkHandler{Service: s}

	r := mux.NewRouter()

	r.HandleFunc("/api/links", h.CreateLink).Methods("POST")
	r.HandleFunc("/api/links/{id}", h.GetLink).Methods("GET")

	fmt.Println("http://localhost:8080")

	http.ListenAndServe(":8080", r)
}
