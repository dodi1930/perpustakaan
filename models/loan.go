package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	ID         uint       `gorm:"primaryKey"`
	UserID     uint       `gorm:"not null"`
	BookID     uint       `gorm:"not null"`
	BorrowDate time.Time  `gorm:"not null"`
	DueDate    time.Time  `gorm:"not null"`
	ReturnDate *time.Time // Bisa null jika belum dikembalikan
	Fine       float64    `gorm:"default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
