package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load variabel dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Tidak bisa memuat file .env, pastikan file ada.")
	}

	// Debugging: Cetak variabel environment
	fmt.Println("🔍 ENV Loaded:")
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))

	// Format string koneksi database
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
		//os.Getenv("DB_TIMEZONE"),
	)

	// Debugging: Cetak DSN
	fmt.Println("🔍 DSN yang digunakan:", dsn)

	// Buka koneksi database dengan GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal koneksi ke database: %v", err)
	}

	DB = db

	fmt.Println("✅ Database berhasil terkoneksi dengan GORM!")

}
