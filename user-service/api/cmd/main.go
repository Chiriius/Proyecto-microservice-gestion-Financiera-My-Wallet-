package main

import (
	"context"
	"my_wallet/api/server"
	"os"

	"github.com/joho/godotenv"
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

	err := godotenv.Load("../../../.env")
	if err != nil {
		logrus.Panic("Layer: main ", "Error al cargar el archivo .env:", err)
	}

	enviromentsVariables := map[string]string{
		"SERVER_PORT_HTTP": os.Getenv("SERVER_PORT_HTTP"),
		"DB_URL":           os.Getenv("DB_URL"),
		"SERVER_PORT_GRPC": os.Getenv("SERVER_PORT_GRPC"),
	}

	entries, err := os.ReadDir("./")
	if err != nil {
		logrus.Fatal(err)
	}

	for _, e := range entries {
		logrus.Info(e.Name())
	}

	httpAddr := enviromentsVariables["SERVER_PORT_HTTP"]
	grpc := enviromentsVariables["SERVER_PORT_GRPC"]
	dburl := enviromentsVariables["DB_URL"]
	srv, err := server.New(logger, httpAddr, grpc, dburl, ctx)
	if err != nil {
		logger.Panic("Layer: main ", "Failed to create server:", err)
	}

	defer srv.Close()
	if err := srv.Start(); err != nil {
		logger.Error(err)
	}

}
