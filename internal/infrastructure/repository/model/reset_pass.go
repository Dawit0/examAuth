package model

import "time"

type PasswordResetModel struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index"`
	Phone     string    `gorm:"size:32;index"`
	OTP       string    `gorm:"size:10"`
	ExpiresAt time.Time `gorm:"index"`
	Used      bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
