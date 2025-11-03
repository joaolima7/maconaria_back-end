package main

import (
	"log"

	"github.com/joaolima7/maconaria_back-end/config"
	"github.com/joaolima7/maconaria_back-end/internal/infra/di"
)

func main() {
	// 1. Carrega configurações do .env
	cfg, err := config.LoadConfig("config/.env")
	if err != nil {
		log.Fatalf("❌ Erro ao carregar configurações: %v", err)
	}

	log.Println("✅ Configurações carregadas com sucesso")

	// 2. Wire injeta TODAS as dependências automaticamente
	_, cleanup, err := di.InitializeUserUseCases(cfg)
	if err != nil {
		log.Fatalf("❌ Erro ao inicializar dependências: %v", err)
	}
	defer cleanup()

	log.Println("✅ Banco de dados conectado")
	log.Println("✅ Dependências injetadas")

}
