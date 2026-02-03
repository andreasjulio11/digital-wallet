package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string
	Balanced  float64 `gorm:"type:decimal;default:0"`
	State     int     `gorm:"type:integer;defau:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
