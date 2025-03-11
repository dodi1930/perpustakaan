package models

import "time"

//type FineSetting struct {
//	ID         uint    `gorm:"primaryKey"`
//	FinePerDay float64 `gorm:"not null"`
//	CreatedAt  time.Time
//	UpdatedAt  time.Time
//}

type BorrowedBook struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"` // Relasi ke User
	BookID     uint      `gorm:"not null"` // Relasi ke Buku
	BorrowDate time.Time `gorm:"not null"`
	DueDate    time.Time `gorm:"not null"`
	ReturnDate *time.Time
	FineAmount float64 `gorm:"default:0"` // Denda per transaksi peminjaman
}
