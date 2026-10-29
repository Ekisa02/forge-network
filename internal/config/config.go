package config

import (
	"os"
)

// Config holds environment-driven configuration.
type Config struct {
	Port        string
	DatabaseURL string
	RPCURL      string
}

// LoadFromEnv reads basic configuration from the environment with sane defaults.
func LoadFromEnv() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	db := os.Getenv("DATABASE_URL")
	if db == "" {
		// Default to sqlite file for easy local dev
		db = "file:forge.db"
	}
	rpc := os.Getenv("RPC_URL")

	return &Config{
		Port:        port,
		DatabaseURL: db,
		RPCURL:      rpc,
	}
}
