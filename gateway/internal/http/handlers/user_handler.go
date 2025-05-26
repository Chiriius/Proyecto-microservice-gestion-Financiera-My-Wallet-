package handlers

import (
	"gateway-mywallet/internal/grpc/clients"
	"gateway-mywallet/internal/grpc/proto/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	user clients.UserClient
}

func NewUserHandler(user clients.UserClient) *UserHandler {
	return &UserHandler{user: user}
}
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Dni      int64  `json:"dni"`
		TypeDNi  string `json:"typeDni"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Address  string `json:"address"`
		Phone    int64  `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requ := &user.CreateUserRequest{
		Dni:      req.Dni,
		TypeDNI:  req.TypeDNi,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Phone:    req.Phone,
	}

	log.Printf("Datos recibidos en HTTP: %+v", req)
	log.Printf("Datos enviados a gRPC: %+v", requ)

	resp, err := h.user.CreateUser(c, requ)
	if err != nil {
		log.Printf("Error al llamar a CreateUser en el servidor gRPC: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
