package main

import (
	"log"
	"net/http"

	"github.com/HuaiyuZhang/benben-ml/internal/config"
	"github.com/HuaiyuZhang/benben-ml/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create new server instance
	srv := server.NewServer(cfg)

	// Initialize routes
	srv.InitializeRoutes()

	// Start the server
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, srv.Router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
