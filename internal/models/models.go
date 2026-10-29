package models

import "time"

// Wallet represents a user or system wallet. NOTE: private_key_encrypted is for demo only.
type Wallet struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	Address             string    `gorm:"uniqueIndex;size:128" json:"address"`
	PublicKey           string    `gorm:"type:text" json:"public_key"`
	PrivateKeyEncrypted string    `gorm:"type:text" json:"private_key_encrypted,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
}

// Transaction is a simple record of relayed transactions.
type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TxHash    string    `gorm:"uniqueIndex;size:128" json:"tx_hash"`
	From      string    `json:"from"`
	To        string    `json:"to,omitempty"`
	Data      string    `gorm:"type:jsonb" json:"data,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
