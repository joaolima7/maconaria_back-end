package migrations

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrationService struct {
	db             *sql.DB
	migrationsPath string
}

func NewMigrationService(db *sql.DB, migrationsPath string) *MigrationService {
	return &MigrationService{
		db:             db,
		migrationsPath: migrationsPath,
	}
}

func (m *MigrationService) getMigrate() (*migrate.Migrate, error) {
	driver, err := mysql.WithInstance(m.db, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao criar driver mysql: %w", err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", m.migrationsPath),
		"mysql",
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar instância de migration: %w", err)
	}

	return migration, nil
}

func (m *MigrationService) Up() error {
	migration, err := m.getMigrate()
	if err != nil {
		return err
	}
	defer migration.Close()

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("erro ao executar migrations up: %w", err)
	}

	version, dirty, err := migration.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return fmt.Errorf("erro ao obter versão: %w", err)
	}

	if errors.Is(err, migrate.ErrNilVersion) {
		log.Println("✅ Nenhuma migration foi executada ainda")
	} else {
		log.Printf("✅ Migrations executadas com sucesso! Versão atual: %d (dirty: %v)", version, dirty)
	}

	return nil
}

func (m *MigrationService) Down() error {
	migration, err := m.getMigrate()
	if err != nil {
		return err
	}
	defer migration.Close()

	if err := migration.Steps(-1); err != nil {
		return fmt.Errorf("erro ao reverter migration: %w", err)
	}

	version, dirty, err := migration.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return fmt.Errorf("erro ao obter versão: %w", err)
	}

	if errors.Is(err, migrate.ErrNilVersion) {
		log.Println("✅ Todas as migrations foram revertidas")
	} else {
		log.Printf("✅ Migration revertida com sucesso! Versão atual: %d (dirty: %v)", version, dirty)
	}

	return nil
}

func (m *MigrationService) DownAll() error {
	migration, err := m.getMigrate()
	if err != nil {
		return err
	}
	defer migration.Close()

	if err := migration.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("erro ao reverter todas migrations: %w", err)
	}

	log.Println("✅ Todas as migrations foram revertidas com sucesso!")
	return nil
}

func (m *MigrationService) Force(version int) error {
	migration, err := m.getMigrate()
	if err != nil {
		return err
	}
	defer migration.Close()

	if err := migration.Force(version); err != nil {
		return fmt.Errorf("erro ao forçar versão %d: %w", version, err)
	}

	log.Printf("✅ Versão forçada para %d com sucesso!", version)
	return nil
}

func (m *MigrationService) Version() error {
	migration, err := m.getMigrate()
	if err != nil {
		return err
	}
	defer migration.Close()

	version, dirty, err := migration.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return fmt.Errorf("erro ao obter versão: %w", err)
	}

	if errors.Is(err, migrate.ErrNilVersion) {
		log.Println("ℹ️  Nenhuma migration foi executada ainda")
	} else {
		status := "limpo"
		if dirty {
			status = "dirty (requer correção)"
		}
		log.Printf("ℹ️  Versão atual: %d (%s)", version, status)
	}

	return nil
}

func (m *MigrationService) Migrate(targetVersion uint) error {
	migration, err := m.getMigrate()
	if err != nil {
		return err
	}
	defer migration.Close()

	if err := migration.Migrate(targetVersion); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("erro ao migrar para versão %d: %w", targetVersion, err)
	}

	log.Printf("✅ Migrado para versão %d com sucesso!", targetVersion)
	return nil
}
