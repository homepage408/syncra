# Syncra

Syncra adalah backend service untuk sistem manajemen workspace multi-tenant yang dibangun dengan arsitektur modular dan clean architecture. Repo ini berfokus pada backend API untuk autentikasi, manajemen post, dan pondasi untuk GraphQL serta event-driven architecture.

## 🚀 Fitur Utama

- Auth REST API:
  - `POST /api/v1/auth/login`
  - `POST /api/v1/auth/register`
  - `POST /api/v1/auth/refresh`
- Health check sederhana:
  - `GET /` atau `GET /health`
- Modular clean architecture dengan pemisahan:
  - domain
  - application/usecase
  - infrastructure/persistence
  - interface/rest
  - bootstrap
- Persistence menggunakan PostgreSQL dan SQLC
- JWT authentication foundation
- GraphQL starter setup dengan gqlgen (placeholder)
- Docker Compose untuk PostgreSQL dan Redis

## 📁 Struktur Proyek

- `cmd/api/main.go` - entrypoint aplikasi
- `internal/bootstrap/` - inisialisasi aplikasi, database, router, dan GraphQL setup
- `internal/domains/` - domain logic terpisah per fitur
- `db/sqlc/` - package query SQLC
- `config/` - konfigurasi environment
- `pkg/` - utilitas shared, logger, JWT, password, response
- `db/migrations/` - skrip migrasi database
- `docs/` - dokumentasi arsitektur

## 🧰 Teknologi

- Go 1.25
- Gin
- gqlgen
- SQLC
- PostgreSQL
- Redis
- JWT

## ⚙️ Persyaratan

- Go 1.25+
- Docker dan Docker Compose
- PostgreSQL (via Docker Compose)

## 💡 Konfigurasi

Aplikasi membaca konfigurasi dari environment variable.

Default environment:

```bash
export DATABASE_DSN=""
export LOG_LEVEL="info"
```

Jika tidak diset, aplikasi akan menggunakan nilai default tersebut.

> Catatan: Port HTTP pada `internal/bootstrap/router_http.go` saat ini dikodekan sebagai `:4500`.

## 🏃 Menjalankan Aplikasi

1. Jalankan database dengan Docker Compose:

```bash
docker compose up -d
```

2. Set environment variable:

```bash
export DATABASE_DSN=""
export LOG_LEVEL="info"
```

3. Jalankan server:

```bash
go run cmd/api/main.go
```

4. Buka API:

```text
http://localhost:4500
```

## 🔧 Perintah Berguna

- Run aplikasi: `go run cmd/api/main.go`
- Build binary: `go build -o bin/api cmd/api/main.go`
- Jalankan test: `go test ./...`
- Generate SQLC: `sqlc generate`
- Generate gqlgen: `go run github.com/99designs/gqlgen generate`
- Migrasi database naik: `migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/syncra?sslmode=disable" up`
- Migrasi database turun: `migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/syncra?sslmode=disable" down`

## 🧪 Endpoint API

### Auth

- `POST /api/v1/auth/login`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/refresh`

### Health

- `GET /`
- `GET /health`

## 📌 Catatan Implementasi

- Endpoint auth dan post saat ini mengembalikan response placeholder.
- GraphQL sudah disiapkan di bootstrap, tetapi resolver dan schema masih dalam tahap pengembangan.
- Konfigurasi `config/app.yaml` disediakan sebagai referensi struktur, namun aplikasi saat ini menggunakan environment variable untuk konfigurasi runtime.

## 📈 Roadmap / Pengembangan Selanjutnya

- Implementasi usecase auth lengkap dan JWT validation
- Konsumsi repository SQLC untuk CRUD post
- Tambahkan middleware autentikasi untuk route post
- Lengkapi GraphQL schema dan resolver
- Tambahkan unit test dan integration test
- Tambahkan fitur multi-tenant workspace, role, permission, dan event-driven notification

## 📄 Referensi

- `docs/architecture/syncra-overview.md`
- `internal/bootstrap/*`
- `db/migrations/*`
- `graph/schema/schema.graphqls`
