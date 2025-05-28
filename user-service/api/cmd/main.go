package main

import (
	"context"
	"my_wallet/api/config"
	"my_wallet/api/server"

	"github.com/sirupsen/logrus"
)

// @title My Wallet API
// @version 1.0
// @description This is a sample server for a wallet API.
// @host localhost:8081
// @BasePath /
func main() {
	ctx := context.Background()
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{})

	cfg := config.Load()

	httpAddr := cfg.Secrets["SERVER_PORT_HTTP"]
	grpcAddr := cfg.Secrets["SERVER_PORT_GRPC"]
	dbURL := cfg.Secrets["DB_URL"]

	if httpAddr == "" || grpcAddr == "" || dbURL == "" {
		logger.Panic("Layer: main ", "Faltan variables de entorno críticas en Vault")
	}

	logger.Infof("HTTP Address: %s", httpAddr)
	logger.Infof("gRPC Address: %s", grpcAddr)
	logger.Infof("Database URL: %s", dbURL)

	srv, err := server.New(logger, httpAddr, grpcAddr, dbURL, ctx)
	if err != nil {
		logger.Panic("Layer: main ", "Failed to create server:", err)
	}

	defer srv.Close()

	if err := srv.Start(); err != nil {
		logger.Error(err)
	}

}
