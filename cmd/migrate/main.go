package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joaolima7/maconaria_back-end/config"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/migrations"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Carrega configurações
	cfg, err := config.LoadConfig("config/.env")
	if err != nil {
		log.Fatalf("❌ Erro ao carregar configurações: %v", err)
	}

	// Conecta ao banco
	db, err := sql.Open(cfg.DBDriver, cfg.GetDSN())
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Erro ao pingar banco: %v", err)
	}

	// Cria serviço de migrations
	migrationService := migrations.NewMigrationService(db, "sql/migrations")

	// Executa comando
	command := os.Args[1]

	switch command {
	case "up":
		if err := migrationService.Up(); err != nil {
			log.Fatalf("❌ Erro: %v", err)
		}

	case "down":
		if err := migrationService.Down(); err != nil {
			log.Fatalf("❌ Erro: %v", err)
		}

	case "down-all":
		fmt.Print("⚠️  Tem certeza que deseja reverter TODAS as migrations? (y/N): ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "y" && confirm != "Y" {
			log.Println("❌ Operação cancelada")
			return
		}
		if err := migrationService.DownAll(); err != nil {
			log.Fatalf("❌ Erro: %v", err)
		}

	case "force":
		if len(os.Args) < 3 {
			log.Fatal("❌ Versão não especificada. Use: migrate force <version>")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("❌ Versão inválida: %v", err)
		}
		if err := migrationService.Force(version); err != nil {
			log.Fatalf("❌ Erro: %v", err)
		}

	case "version":
		if err := migrationService.Version(); err != nil {
			log.Fatalf("❌ Erro: %v", err)
		}

	case "goto":
		if len(os.Args) < 3 {
			log.Fatal("❌ Versão não especificada. Use: migrate goto <version>")
		}
		version, err := strconv.ParseUint(os.Args[2], 10, 32)
		if err != nil {
			log.Fatalf("❌ Versão inválida: %v", err)
		}
		if err := migrationService.Migrate(uint(version)); err != nil {
			log.Fatalf("❌ Erro: %v", err)
		}

	default:
		fmt.Printf("❌ Comando desconhecido: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("📦 Gerenciador de Migrations")
	fmt.Println("\nUso: go run cmd/migrate/main.go <comando> [argumentos]")
	fmt.Println("\nComandos disponíveis:")
	fmt.Println("  up          - Executa todas as migrations pendentes")
	fmt.Println("  down        - Reverte a última migration")
	fmt.Println("  down-all    - Reverte TODAS as migrations (com confirmação)")
	fmt.Println("  version     - Mostra a versão atual das migrations")
	fmt.Println("  force <n>   - Força uma versão específica (útil para corrigir estado dirty)")
	fmt.Println("  goto <n>    - Migra para uma versão específica")
	fmt.Println("\nExemplos:")
	fmt.Println("  go run cmd/migrate/main.go up")
	fmt.Println("  go run cmd/migrate/main.go down")
	fmt.Println("  go run cmd/migrate/main.go version")
	fmt.Println("  go run cmd/migrate/main.go force 2")
	fmt.Println("  go run cmd/migrate/main.go goto 3")
}
