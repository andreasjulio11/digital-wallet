package models

import "time"

type DetailTransaction struct {
	ID              int     `gorm:"primaryKey"`
	UserID          int     `gorm:"not null"`
	User            User    `gorm:"foreingKey:UserID"`
	TransactionType string  `gorm:"type:varchar(100);not null"`
	Amount          float64 `gorm:"type:decimal(16,2);not null"`
	CreatedAt       time.Time
}
