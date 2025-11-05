package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

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

	// Ping com timeout para evitar bloqueios longos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := database.PingContext(ctx); err != nil {
		database.Close()
		return nil, fmt.Errorf("erro ao conectar no banco: %w", err)
	}

	database.SetMaxOpenConns(25)
	database.SetMaxIdleConns(10)

	// Mantenha ConnMaxIdleTime menor que o wait_timeout do MySQL (ex.: 30s)
	database.SetConnMaxIdleTime(30 * time.Second)
	database.SetConnMaxLifetime(5 * time.Minute)

	if cfg.AutoMigrate {
		log.Println("🔄 Executando migrations automaticamente...")

		migrationDB, err := sql.Open(cfg.DBDriver, cfg.GetDSN())
		if err != nil {
			log.Printf("⚠️  Aviso: erro ao criar conexão para migrations: %v", err)
		} else {
			// garante fechamento da conexão de migrations
			defer migrationDB.Close()

			migrationService := migrations.NewMigrationService(migrationDB, "sql/migrations")
			if err := migrationService.Up(); err != nil {
				log.Printf("⚠️  Aviso: erro ao executar migrations: %v", err)
			}
		}
	}

	// Keep-alive menos agressivo e com timeout de contexto
	go keepAlive(database)

	return database, nil
}

func keepAlive(db *sql.DB) {
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		if err := db.Ping(); err != nil {
			log.Printf("⚠️  Keep-alive ping falhou: %v", err)
		}
	}
}

func ProvideQueries(database *sql.DB) *db.Queries {
	return db.New(database)
}
