package handlers

import (
	"gateway-mywallet/internal/grpc/clients"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	client clients.AccountClient
}

func NewAccountHandler(client clients.AccountClient) *AccountHandler {
	return &AccountHandler{client: client}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req struct {
		UserID      string  `json:"userId"`
		Balance     float64 `json:"balance"`
		AccountName string  `json:"accountName"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Datos recibidos en HTTP: %+v", req)

	resp, err := h.client.CreateAccount(c, req.UserID, req.Balance, req.AccountName)
	if err != nil {
		log.Printf("Error al llamar a CreateAccount: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
