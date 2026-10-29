package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yourusername/forge-network/internal/services"
)

type WalletHandler struct {
	svc *services.WalletService
}

func NewWalletHandler(s *services.WalletService) *WalletHandler { return &WalletHandler{svc: s} }

func (h *WalletHandler) Create(c *gin.Context) {
	w, err := h.svc.CreateWallet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return only non-sensitive metadata in responses
	c.JSON(http.StatusCreated, gin.H{"address": w.Address, "public_key": w.PublicKey})
}

func (h *WalletHandler) Get(c *gin.Context) {
	addr := c.Param("address")
	w, err := h.svc.GetByAddress(addr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if w == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"address": w.Address, "public_key": w.PublicKey})
}
