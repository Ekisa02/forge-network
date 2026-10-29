package services

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/yourusername/forge-network/internal/models"
	"github.com/yourusername/forge-network/internal/repository"
)

// WalletService handles wallet-related business logic.
type WalletService struct {
	repo *repository.WalletRepo
}

func NewWalletService(r *repository.WalletRepo) *WalletService { return &WalletService{repo: r} }

// CreateWallet generates a new ECDSA keypair and persists metadata. For demo purposes the private key
// is stored in plaintext hex in PrivateKeyEncrypted — replace with proper encryption in production.
func (s *WalletService) CreateWallet() (*models.Wallet, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	addr := crypto.PubkeyToAddress(priv.PublicKey).Hex()
	pubBytes := crypto.FromECDSAPub(&priv.PublicKey)
	privBytes := crypto.FromECDSA(priv)

	w := &models.Wallet{
		Address:             addr,
		PublicKey:           fmt.Sprintf("0x%s", hex.EncodeToString(pubBytes)),
		PrivateKeyEncrypted: hex.EncodeToString(privBytes), // DEMOONLY
	}

	if err := s.repo.Create(w); err != nil {
		return nil, err
	}
	return w, nil
}

func (s *WalletService) GetByAddress(addr string) (*models.Wallet, error) {
	return s.repo.GetByAddress(addr)
}
