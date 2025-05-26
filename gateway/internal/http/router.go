package http

import (
	"gateway-mywallet/internal/grpc/clients"
	"gateway-mywallet/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(accountClient clients.AccountClient, userClient clients.UserClient) *gin.Engine {
	router := gin.Default()

	accountHandler := handlers.NewAccountHandler(accountClient)
	userHandler := handlers.NewUserHandler(userClient)
	router.POST("/accounts", accountHandler.CreateAccount)
	router.POST("/users", userHandler.CreateUser)

	return router
}
