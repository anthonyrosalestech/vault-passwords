// models/user.go
package models

import "time"

type User struct {
	ID             string    `gorm:"primaryKey"`
	Email          string    `gorm:"uniqueIndex;size:255"`
	Username       string    `gorm:"size:100"`
	PasswordHash   string    `gorm:"type:text"`
	PasswordSalt   string    `gorm:"type:text"`
	Is2FAEnabled   bool
	TwoFASecret    string    `gorm:"type:text"`
	TwoFAVerified  bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}