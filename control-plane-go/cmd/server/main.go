package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sunilganta-dev/agentic-ml-control-plane/control-plane-go/internal/config"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := config.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	log.Println("Connected to database")

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthHandler)

	log.Printf("Control Plane starting on port %s...", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}

	_ = db
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
