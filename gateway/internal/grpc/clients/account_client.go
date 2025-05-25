package clients

import (
	"context"
	account "gateway-mywallet/internal/grpc/proto"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type AccountClient interface {
	CreateAccount(ctx context.Context, userId string, balance float64, accountName string) (*account.CreateAccountResponse, error)
}

type accountClient struct {
	client account.AccountServiceClient
}

func NewAccountClient(address string) AccountClient {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor gRPC: %v", err)
	}
	return &accountClient{client: account.NewAccountServiceClient(conn)}
}

func (c *accountClient) CreateAccount(ctx context.Context, userId string, balance float64, accountName string) (*account.CreateAccountResponse, error) {
	req := &account.CreateAccountRequest{
		UserId:      userId,
		Balance:     float32(balance),
		AccountName: accountName,
	}

	resp, err := c.client.CreateAccount(ctx, req)
	if err != nil {
		log.Printf("Error al llamar a CreateAccount en el servidor gRPC: %v", err)
		return nil, err
	}

	return resp, nil
}
