ifneq (,$(wildcard config/.env))
include config/.env
export
endif

DB_USER     ?= root
DB_PASSWORD ?= root
DB_NAME     ?= maconaria_db
DB_HOST     ?= 127.0.0.1
DB_PORT     ?= 3306

MIGRATE_URL ?= mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

# Criar nova migration
createMigration:
    migrate create -ext=sql -dir=sql/migrations -seq $(MIGRATION_NAME)

# Executar migrations (via CLI personalizado)
migrate-up:
    @go run cmd/migrate/main.go up

migrate-down:
    @go run cmd/migrate/main.go down

migrate-down-all:
    @go run cmd/migrate/main.go down-all

migrate-version:
    @go run cmd/migrate/main.go version

migrate-force:
    @go run cmd/migrate/main.go force $(VERSION)

migrate-goto:
    @go run cmd/migrate/main.go goto $(VERSION)

# Executar aplicação
run:
    @go run cmd/main.go

# Build
build:
    @go build -o bin/server cmd/main.go

# Gerar código sqlc
sqlc-generate:
    @sqlc generate

# Gerar código wire
wire-generate:
    @cd internal/infra/di && wire

# Setup completo (wire + sqlc)
setup:
    @make wire-generate
    @make sqlc-generate

# Instalar dependências
deps:
    @go mod download
    @go mod tidy

.PHONY: createMigration migrate-up migrate-down migrate-down-all migrate-version migrate-force migrate-goto run build sqlc-generate wire-generate setup deps