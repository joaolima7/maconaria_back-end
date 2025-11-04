package main

import (
	"log"

	"github.com/joaolima7/maconaria_back-end/config"
	"github.com/joaolima7/maconaria_back-end/internal/infra/di"
)

func main() {
	cfg, err := config.LoadConfig("config/.env")
	if err != nil {
		log.Fatalf("❌ Erro ao carregar configurações: %v", err)
	}

	log.Println("✅ Configurações carregadas")

	server, cleanup, err := di.InitializeServer(cfg)
	if err != nil {
		log.Fatalf("❌ Erro ao inicializar servidor: %v", err)
	}
	defer cleanup()

	log.Println("✅ Banco de dados conectado")
	log.Println("✅ Dependências injetadas")

	if err := server.Start(); err != nil {
		log.Fatalf("❌ Erro no servidor: %v", err)
	}
}
