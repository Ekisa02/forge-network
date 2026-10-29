package server

import (
	"github.com/gin-gonic/gin"

	"github.com/yourusername/forge-network/internal/config"
	"github.com/yourusername/forge-network/internal/db"
	"github.com/yourusername/forge-network/internal/handlers"
	"github.com/yourusername/forge-network/internal/repository"
	"github.com/yourusername/forge-network/internal/services"
)

type Server struct {
	r *gin.Engine
	db *db.DB
	cfg *config.Config
}

func NewServer(d *db.DB, cfg *config.Config) *Server {
	r := gin.Default()
	// Initialize layers
	walletRepo := repository.NewWalletRepo(d.DB)
	walletSvc := services.NewWalletService(walletRepo)
	walletHandler := handlers.NewWalletHandler(walletSvc)

	// Routes
	r.POST("/wallets", walletHandler.Create)
	r.GET("/wallets/:address", walletHandler.Get)

	return &Server{r: r, db: d, cfg: cfg}
}

func (s *Server) Run(addr string) error { return s.r.Run(addr) }
