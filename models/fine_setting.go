package models

import (
	"gorm.io/gorm"
)

// FineSetting menyimpan nilai denda per hari
type FineSetting struct {
	gorm.Model
	FinePerDay float64 `json:"fine_per_day"`
}
