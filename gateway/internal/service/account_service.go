package service

import (
	"context"
	"gateway-mywallet/internal/grpc/clients"
	account "gateway-mywallet/internal/grpc/proto"
)

type AccountService interface {
	CreateAccount(ctx context.Context, userId string, balance float64, accountName string) (*account.CreateAccountResponse, error)
}

type accountService struct {
	client clients.AccountClient
}

func NewAccountUsecase(client clients.AccountClient) AccountService {
	return &accountService{client: client}
}

func (u *accountService) CreateAccount(ctx context.Context, userId string, balance float64, accountName string) (*account.CreateAccountResponse, error) {
	return u.client.CreateAccount(ctx, userId, balance, accountName)
}
