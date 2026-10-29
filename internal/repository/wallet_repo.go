package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/yourusername/forge-network/internal/models"
)

type WalletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) *WalletRepo { return &WalletRepo{db: db} }

func (r *WalletRepo) Create(w *models.Wallet) error {
	return r.db.Create(w).Error
}

func (r *WalletRepo) GetByAddress(address string) (*models.Wallet, error) {
	var w models.Wallet
	if err := r.db.Where("address = ?", address).First(&w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &w, nil
}
