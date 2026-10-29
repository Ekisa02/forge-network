# Forge Network

Forge Network is a modular Web3 infrastructure toolkit written in Go using Gin and GORM. It provides a minimal, secure foundation for wallet management, transaction relaying, and an indexable store for blockchain-related events — intended as a starting point for building production-grade infrastructure.

Key features
- Wallet lifecycle management (create, lookup) — demo-only private key handling (do NOT store plaintext keys in production)
- Simple transaction relay stub (extensible to multiple chains / providers)
- GORM data layer with migrations and easy switching between Postgres and SQLite for tests
- Gin HTTP API with clear modular structure for services, repositories, and handlers

Architecture

- cmd/server: application entrypoint
- internal/config: environment-driven configuration loader
- internal/db: database connection and migration helpers
- internal/models: GORM models (Wallet, Transaction)
- internal/repository: data access objects (repositories) for models
- internal/services: business logic (wallet creation, relayer stub)
- internal/handlers: HTTP handlers for endpoints
- migrations: SQL migration (Postgres)

Installation

Prerequisites
- Go 1.20+
- Docker & Docker Compose (optional, for running Postgres locally)

Quick start (local, using Docker Postgres):

1. Copy .env.example to .env and edit values if needed.

2. Start Postgres (Docker):

   docker-compose up -d

3. Build and run:

   make build
   ./bin/forge-server

4. API

   POST  /wallets         -> create a new wallet
   GET   /wallets/:addr   -> fetch wallet metadata

Example (create wallet):

curl -X POST http://localhost:8080/wallets -H "Content-Type: application/json"

Response:
{
  "address": "0x..."
}

Development & tests

- Run unit tests:

  make test

- Run with an in-memory SQLite DB for local testing by setting DATABASE_URL to "file::memory:?cache=shared" (or using the default .env)

Security note

This project stores private keys in plaintext only for demonstration and local development. DO NOT store keys in plaintext in production — use HSMs, KMS, or an encryption-at-rest strategy and rotate keys frequently.

Contributing

Contributions welcome. Follow Go and module idioms and open PRs against the main branch. Include tests for new functionality.
