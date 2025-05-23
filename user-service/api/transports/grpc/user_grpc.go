package transport_grpc

import (
	"context"
	"fmt"
	"my_wallet/api/endpoints"

	pb "my_wallet/api/proto"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/sirupsen/logrus"
)

type gRPCServer struct {
	createUser     gt.Handler
	loginUser      gt.Handler
	getUser        gt.Handler
	deleteUser     gt.Handler
	updateUser     gt.Handler
	softDeleteUser gt.Handler
	pb.UnimplementedUserServiceServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger logrus.FieldLogger) pb.UserServiceServer {
	return &gRPCServer{
		createUser: gt.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),
		loginUser: gt.NewServer(
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

func (s *gRPCServer) mustEmbedUnimplementedUserServiceServer() {}

func (s *gRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateUserResponse), nil
}
func (s *gRPCServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	_, resp, err := s.loginUser.ServeGRPC(ctx, req)
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
	resp, ok := response.(endpoints.CreateUserResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response type: expected CreateUserResponse, got %T", response)
	}

	return &pb.CreateUserResponse{
		Id:    resp.ID,
		Token: resp.Token,
		Error: resp.Err,
	}, nil
}
func decodeLoginUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoginUserRequest)
	return endpoints.LoginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
func encodeLoginUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.LoginUserResponse)
	return &pb.LoginUserResponse{Token: resp.Token, Error: resp.Err}, nil
}
func decodeGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserRequest)
	return endpoints.GetUserRequest{
		ID: req.Id,
	}, nil
}
func encodeGetUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GetUserResponse)
	user := &pb.User{
		Id:      resp.User.ID,
		Dni:     int64(resp.User.DNI),
		TypeDNI: resp.User.TypeDNI,
		Name:    resp.User.Name,
		Email:   resp.User.Email,
		Address: resp.User.Address,
		Phone:   int64(resp.User.Phone),
	}
	return &pb.GetUserResponse{
		User:  user,
		Error: resp.Err,
	}, nil
}
func decodeDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteUserRequest)
	return endpoints.DeleteUserRequest{
		ID: req.Id,
	}, nil
}
func encodeDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.DeleteUserResponse)
	return &pb.DeleteUserResponse{Error: resp.Err}, nil
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
	resp := response.(endpoints.UpdateUserREsponse)
	user := &pb.User{
		Id:      resp.User.ID,
		Dni:     int64(resp.User.DNI),
		TypeDNI: resp.User.TypeDNI,
		Name:    resp.User.Name,
		Email:   resp.User.Email,
		Address: resp.User.Address,
		Phone:   int64(resp.User.Phone),
	}
	return &pb.UpdateUserREsponse{User: user, Error: resp.Err}, nil
}
func decodeSoftDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SoftDeleteUserRequest)
	return endpoints.SoftDeleteUserRequest{
		ID: req.Id,
	}, nil
}
func encodeSoftDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.SoftDeleteUserResponse)
	return &pb.SoftDeleteUserResponse{Error: resp.Err}, nil
}
