package services_test

import (
	"testing"

	"github.com/yourusername/forge-network/internal/models"
	"github.com/yourusername/forge-network/internal/repository"
	"github.com/yourusername/forge-network/internal/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&models.Wallet{}); err != nil {
		t.Fatalf("auto migrate failed: %v", err)
	}
	return db
}

func TestCreateWallet(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewWalletRepo(db)
	svc := services.NewWalletService(repo)

	w, err := svc.CreateWallet()
	if err != nil {
		t.Fatalf("CreateWallet failed: %v", err)
	}
	if w.Address == "" {
		t.Fatalf("expected address, got empty")
	}

	// verify persisted
	got, err := repo.GetByAddress(w.Address)
	if err != nil {
		t.Fatalf("repo GetByAddress failed: %v", err)
	}
	if got == nil || got.Address != w.Address {
		t.Fatalf("expected to find wallet, got nil or mismatch")
	}
}
