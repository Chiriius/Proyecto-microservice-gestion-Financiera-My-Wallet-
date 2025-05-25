package main

import (
	"gateway-mywallet/config"
	"gateway-mywallet/internal/grpc/clients"
	"gateway-mywallet/internal/http"
	"log"
)

func main() {
	// Cargar configuración
	cfg := config.Load()

	// Inicializar clientes gRPC
	accountClient := clients.NewAccountClient(cfg.AccountServiceAddress)

	// Iniciar servidor HTTP
	router := http.NewRouter(accountClient)
	log.Printf("Servidor HTTP ejecutándose en el puerto %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
