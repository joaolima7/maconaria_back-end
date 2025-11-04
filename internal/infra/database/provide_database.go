package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joaolima7/maconaria_back-end/config"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/migrations"
)

func ProvideDatabase(cfg *config.Config) (*sql.DB, error) {
	database, err := sql.Open(cfg.DBDriver, cfg.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
	}

	if err := database.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco: %w", err)
	}

	database.SetMaxOpenConns(25)
	database.SetMaxIdleConns(5)

	if cfg.AutoMigrate {
		log.Println("🔄 Executando migrations automaticamente...")

		migrationDB, err := sql.Open(cfg.DBDriver, cfg.GetDSN())
		if err != nil {
			log.Printf("⚠️  Aviso: erro ao criar conexão para migrations: %v", err)
		} else {
			migrationService := migrations.NewMigrationService(migrationDB, "sql/migrations")
			if err := migrationService.Up(); err != nil {
				log.Printf("⚠️  Aviso: erro ao executar migrations: %v", err)
			}

		}
	}

	return database, nil
}

func ProvideQueries(database *sql.DB) *db.Queries {
	return db.New(database)
}
