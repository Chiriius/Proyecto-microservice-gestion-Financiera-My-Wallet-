package main

import (
	"gateway-mywallet/config"
	"gateway-mywallet/internal/grpc/clients"
	"gateway-mywallet/internal/http"
	"log"
)

func main() {

	cfg := config.Load()

	accountAddr := cfg.Secrets["ACCOUNT_SERVICE_ADDRESS"]
	userAddr := cfg.Secrets["USER_SERVICE_ADDRESS"]
	port := cfg.Secrets["PORT_GATEWAY"]

	accountClient := clients.NewAccountClient(accountAddr)
	userClient := clients.NewUserClient(userAddr)

	router := http.NewRouter(accountClient, userClient)
	log.Printf("Servidor HTTP ejecutándose en el puerto %s", port)
	router.Run(":" + port)
}
