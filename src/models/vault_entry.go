// models/vault_entry.go
package models

import "time"

type VaultEntry struct {
	ID         string `gorm:"primaryKey"`
	UserID     string
	Name       string
	Username   string
	PasswordEnc string
	WebsiteURL string
	Notes      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}