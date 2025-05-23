package transport_grpc

import (
	"context"
	"my_wallet/api/endpoints"

	pb "my_wallet/api/proto"

	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	createUser     gt.Handler
	login          gt.Handler
	getUser        gt.Handler
	deleteUser     gt.Handler
	updateUser     gt.Handler
	softDeleteUser gt.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints) *gRPCServer {
	return &gRPCServer{
		createUser: gt.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),
		login: gt.NewServer(
			endpoints.Login,
			decodeLoginUserRequest,
			encodeLoginUserResponse,
		),
		getUser: gt.NewServer(
			endpoints.GetUser,
			decodeGetUserRequest,
			encodeGetUserResponse,
		),
		deleteUser: gt.NewServer(
			endpoints.DeleteUser,
			decodeDeleteUserRequest,
			encodeDeleteUserResponse,
		),
		updateUser: gt.NewServer(
			endpoints.UpdateUser,
			decodeUpdateRequest,
			encodeUpdateUserResponse,
		),
		softDeleteUser: gt.NewServer(
			endpoints.SoftDeleteUser,
			decodeSoftDeleteUserRequest,
			encodeSoftDeleteUserResponse,
		),
	}
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateUserResponse), nil
}
func (s *gRPCServer) Login(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	_, resp, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.LoginUserResponse), nil
}
func (s *gRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetUserResponse), nil
}
func (s *gRPCServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, resp, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DeleteUserResponse), nil
}

func (s *gRPCServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserREsponse, error) {
	_, resp, err := s.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UpdateUserREsponse), nil
}
func (s *gRPCServer) SoftDeleteUser(ctx context.Context, req *pb.SoftDeleteUserRequest) (*pb.SoftDeleteUserResponse, error) {
	_, resp, err := s.softDeleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SoftDeleteUserResponse), nil
}
func decodeCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return endpoints.CreateUserRequest{
		DNI:      int(req.Dni),
		TypeDNI:  req.TypeDNI,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Phone:    int(req.Phone),
	}, nil
}
func encodeCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.CreateUserResponse)
	return &pb.CreateUserResponse{Id: resp.Id, Token: resp.Token, Error: resp.Error}, nil
}
func decodeLoginUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoginUserRequest)
	return endpoints.LoginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
func encodeLoginUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginUserResponse)
	return &pb.LoginUserResponse{Token: resp.Token, Error: resp.Error}, nil
}
func decodeGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserRequest)
	return endpoints.GetUserRequest{
		ID: req.Id,
	}, nil
}
func encodeGetUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetUserResponse)
	user := &pb.User{
		Id:        resp.User.Id,
		Dni:       int32(resp.User.Dni),
		TypeDNI:   resp.User.TypeDNI,
		Name:      resp.User.Name,
		Email:     resp.User.Email,
		Address:   resp.User.Address,
		Phone:     int32(resp.User.Phone),
		CreatedAt: resp.User.CreatedAt,
	}
	return &pb.GetUserResponse{
		User:  user,
		Error: resp.Error,
	}, nil
}
func decodeDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteUserRequest)
	return endpoints.DeleteUserRequest{
		ID: req.Id,
	}, nil
}
func encodeDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.DeleteUserResponse)
	return &pb.DeleteUserResponse{Error: resp.Error}, nil
}
func decodeUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateUserRequest)
	return endpoints.UpdateUserRequest{
		ID:       req.Id,
		DNI:      int(req.Dni),
		TypeDNI:  req.TypeDNI,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Phone:    int(req.Phone),
	}, nil
}
func encodeUpdateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UpdateUserREsponse)
	user := &pb.User{
		Id:        resp.User.Id,
		Dni:       int32(resp.User.Dni),
		TypeDNI:   resp.User.TypeDNI,
		Name:      resp.User.Name,
		Email:     resp.User.Email,
		Address:   resp.User.Address,
		Phone:     int32(resp.User.Phone),
		CreatedAt: resp.User.CreatedAt,
	}
	return &pb.UpdateUserREsponse{User: user, Error: resp.Error}, nil
}
func decodeSoftDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SoftDeleteUserRequest)
	return endpoints.SoftDeleteUserRequest{
		ID: req.Id,
	}, nil
}
func encodeSoftDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.SoftDeleteUserResponse)
	return &pb.SoftDeleteUserResponse{Error: resp.Error}, nil
}
