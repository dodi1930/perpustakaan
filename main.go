package main

import (
	"fmt"
	"log"
	"os"

	"perpustakaan-x-cgpt/controllers"
	"perpustakaan-x-cgpt/models"
	"perpustakaan-x-cgpt/routers"

	//"perpustakaan-x-cgpt/routers"

	//ithub.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load konfigurasi dari .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Ambil konfigurasi dari .env
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbTimezone := os.Getenv("DB_TIMEZONE")

	// Format DSN untuk koneksi database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimezone,
	)

	// Koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	// Migrasi tabel (pastikan semua tabel dimigrasikan)
	db.AutoMigrate(&models.User{}, &models.Book{}, &models.Loan{}, &models.FineSetting{})

	// Set database di controller agar bisa digunakan di semua fitur
	controllers.SetDatabase(db)

	// Inisialisasi router dengan semua endpoint
	r := routers.SetupRouter()

	//r := gin.Default()
	r.LoadHTMLGlob("templates/dashboard.html") // Pastikan folder templates ada

	// Jalankan server di port 9090
	r.Run(":9090")
}
