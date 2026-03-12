package main

import (
	"github.com/rs/cors"

	"lobster/internal/handler"
	"lobster/internal/repository"
	"lobster/internal/service"

	"fmt"
	"net/http"
	"time"
)

func main() {
	storage := repository.NewLinkStore()
	s := service.NewLinkService(storage)
	h := handler.NewLinkHandler(s)

	r := http.NewServeMux()

	// REGISTER ROUTES
	h.RegisterRoutes(r)

	storage.StartCleanup(2 * time.Minute)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(r)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is starting on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
