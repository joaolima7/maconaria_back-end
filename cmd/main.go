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

	app, err := di.InitializeApp(cfg)
	if err != nil {
		log.Fatalf("❌ Erro ao inicializar aplicação: %v", err)
	}
	defer app.Cleanup() // Agora fecha o banco APÓS o servidor parar

	log.Println("✅ Banco de dados conectado")
	log.Println("✅ Dependências injetadas")

	// Start é bloqueante até receber sinal de interrupção
	if err := app.Server.Start(); err != nil {
		log.Fatalf("❌ Erro no servidor: %v", err)
	}
}
