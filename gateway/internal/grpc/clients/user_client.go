package clients

import (
	"context"
	user "gateway-mywallet/internal/grpc/proto/user"
	"log"

	"google.golang.org/grpc"
)

type UserClient interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error)
}

func NewUserClient(address string) UserClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor gRPC: %v", err)
	}
	return &userClient{user: user.NewUserServiceClient(conn)}
}

type userClient struct {
	user user.UserServiceClient
}

func (c *userClient) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {

	log.Printf("Datos recibidos en gRPC: %+v", req)
	resp, err := c.user.CreateUser(ctx, req)
	if err != nil {
		log.Printf("Error al llamar a CreateUser en el servidor gRPC: %v", err)
		return nil, err
	}

	return resp, nil
}
