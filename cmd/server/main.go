package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yourusername/forge-network/internal/config"
	"github.com/yourusername/forge-network/internal/db"
	"github.com/yourusername/forge-network/internal/server"
)

func main() {
	cfg := config.LoadFromEnv()

	dbConn, err := db.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// Run automatic migrations for demo projects. For production, use proper migrations.
	if err := dbConn.AutoMigrate(); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	srv := server.NewServer(dbConn, cfg)
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("starting server on %s", addr)
	if err := srv.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
	os.Exit(0)
}
