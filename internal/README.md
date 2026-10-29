Internal packages and architecture

- config: simple loader for environment variables and defaults
- db: DB connection wrapper and demo AutoMigrate
- models: GORM models representing domain objects
- repository: persistence layer abstracting GORM queries
- services: business logic (wallet creation, relayer placeholder)
- handlers: HTTP handlers using Gin
- server: bootstraps Gin and wires components together

Use these modules as a pattern to grow the codebase: add a dedicated relayer service, event indexing worker, and RPC provider adapters under internal/providers.
