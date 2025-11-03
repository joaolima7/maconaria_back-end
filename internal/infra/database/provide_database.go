package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joaolima7/maconaria_back-end/config"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
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

	return database, nil
}

func ProvideQueries(database *sql.DB) *db.Queries {
	return db.New(database)
}
