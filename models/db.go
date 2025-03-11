package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Deklarasi variabel global DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL") // Pastikan DATABASE_URL ada di .env
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	DB = db
}
