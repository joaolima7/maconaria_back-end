ifneq (,$(wildcard config/.env))
include config/.env
export
endif

DB_USER     ?= root
DB_PASSWORD ?= root
DB_NAME     ?= maconaria_db
DB_HOST     ?= localhost
DB_PORT     ?= 5432
DB_SSLMODE  ?= disable

MIGRATE_URL ?= postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

createMigration:
	migrate create -ext=sql -dir=sql/migrations -seq $(MIGRATION_NAME)

migrateUp:
	migrate -path sql/migrations -database "$(MIGRATE_URL)" up

migrateDown:
	migrate -path sql/migrations -database "$(MIGRATE_URL)" down

.PHONY: createMigration migrateUp migrateDown
