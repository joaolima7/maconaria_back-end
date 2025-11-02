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

createMigration:
    migrate create -ext=sql -dir=sql/migrations -seq $(MIGRATION_NAME)

migrateUp:
    migrate -path sql/migrations -database "$(MIGRATE_URL)" up

migrateDown:
    migrate -path sql/migrations -database "$(MIGRATE_URL)" down

.PHONY: createMigration migrateUp migrateDown