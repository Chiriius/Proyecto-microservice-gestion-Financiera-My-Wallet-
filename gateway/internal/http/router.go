package http

import (
	"gateway-mywallet/internal/grpc/clients"
	"gateway-mywallet/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(accountClient clients.AccountClient) *gin.Engine {
	router := gin.Default()

	accountHandler := handlers.NewAccountHandler(accountClient)
	router.POST("/accounts", accountHandler.CreateAccount)

	return router
}
