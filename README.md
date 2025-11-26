# Prestasi API (Golang) - Minimal Starter

## Requirements
- Go 1.21
- Docker & Docker Compose

## Setup
1. Copy `.env.example` to `.env` and sesuaikan.
2. Start DB:
   docker-compose up -d
3. Apply migrations (psql):
   docker exec -it <postgres-container> psql -U postgres -d prestasi -f /path/to/migrations/init.sql
   (atau jalankan manual via DB client)
4. Build & run:
   make run  # or `go run ./cmd/server`
5. API:
   POST /api/v1/auth/login
   Protected endpoints: Authorization: Bearer <token>
