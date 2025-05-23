package server

import (
	"context"
	"fmt"
	_ "my_wallet/api/cmd/docs"
	"my_wallet/api/endpoints"
	infraestructure_repository "my_wallet/api/repository/healtcheck"
	repository_user "my_wallet/api/repository/user"
	"net"

	pb "my_wallet/api/proto"
	infraestructure_services "my_wallet/api/services/healtcheck"
	services "my_wallet/api/services/user"
	transport_grpc "my_wallet/api/transports/grpc"
	transports "my_wallet/api/transports/http"

	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	dbMongo  *mongo.Client
	httpMux  *http.ServeMux
	httpAddr string
	grpcSrv  *grpc.Server
	grpcAddr string
	logger   logrus.FieldLogger
}

func New(logger logrus.FieldLogger, httpAddr, grpcAddr, dburl string, ctx context.Context) (*Server, error) {
	db := GetMongoDB(ctx, dburl)

	healtCheckRepository := infraestructure_repository.NewMongoUserREpository(db, logger)
	healtCheckService := infraestructure_services.NewHealtcheckService(ctx, healtCheckRepository, logger)
	userRepository := repository_user.NewMongoUserREpository(db, logger)
	userService := services.NewUserService(userRepository, logger, ctx)
	userEnpoints := endpoints.MakeServerEndpoints(userService, healtCheckService, logger)
	httpHandler := transports.NewHTTPHandler(userEnpoints, logger)
	grpcServer := transport_grpc.NewGRPCServer(userEnpoints, logger)

	httpMux := http.NewServeMux()
	httpMux.Handle("/", httpHandler)
	httpMux.Handle("/swagger/", httpSwagger.WrapHandler)
	httpMux.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/app/api/cmd/docs/swagger.json")
	})

	baseServerGrpc := grpc.NewServer()
	pb.RegisterUserServiceServer(baseServerGrpc, grpcServer)

	return &Server{
		dbMongo:  db,
		httpMux:  httpMux,
		grpcSrv:  baseServerGrpc,
		grpcAddr: grpcAddr,
		httpAddr: httpAddr,
		logger:   logger,
	}, nil
}

func (s *Server) Start() error {
	if s.httpAddr == "" {
		s.logger.Error("Layer:Server Method:Start ", "httpAddr no está definido")
		return fmt.Errorf("httpAddr no está definido")
	}
	if s.httpMux == nil {
		return fmt.Errorf("httpMux no está inicializado")
	}
	if s.logger == nil {
		return fmt.Errorf("logger no está inicializado")
	}

	go func() {
		s.logger.Infoln("Layer:Server Method:Start ", "Starting HTTP server on ", s.httpAddr)
		if err := http.ListenAndServe(s.httpAddr, s.httpMux); err != nil {
			s.logger.Fatalf("Layer:Server Method:Start ", "Failed to start HTTP server: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", s.grpcAddr)
	if err != nil {
		return fmt.Errorf("Failed to start gRPC listener: %v", err)
	}
	s.logger.Infoln("Layer:Server Method:Start ", "Starting gRPC server on ", s.grpcAddr)
	return s.grpcSrv.Serve(lis)
}

func (s *Server) Close() error {
	s.grpcSrv.GracefulStop()
	return s.dbMongo.Disconnect(context.TODO())
}

func GetMongoDB(ctx context.Context, dburl string) *mongo.Client {
	opts := options.Client().ApplyURI(dburl)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Panic("Layel:Server Method:GetMongoDB ", err)
		os.Exit(2)
	}

	if err := client.Ping(ctx, nil); err != nil {
		logrus.Panic("Layel:Server Method:GetMongoDB Ping failed: ", err)
		os.Exit(2)
	}

	return client
}
