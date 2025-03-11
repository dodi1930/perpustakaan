package controllers

import (
	"log"
	"net/http"
	"perpustakaan-x-cgpt/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReturnLoan(c *gin.Context) {
	loanID, err := strconv.Atoi(c.Param("loan_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan ID tidak valid"})
		return
	}

	// Perbarui tanggal pengembalian
	err = repositories.ReturnBook(uint(loanID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengembalikan buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dikembalikan"})
}

// API untuk melihat riwayat peminjaman berdasarkan user_id
func GetLoanHistory(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println("❌ User ID tidak valid:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID tidak valid"})
		return
	}

	loans, err := repositories.GetLoanHistory(uint(userID))
	if err != nil {
		log.Println("❌ Error mengambil riwayat peminjaman:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil riwayat peminjaman"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"history": loans})
}

// API untuk mendapatkan statistik peminjaman (Admin Only)
func GetLoanStatistics(c *gin.Context) {
	stats, err := repositories.GetLoanStatistics()
	if err != nil {
		log.Println("❌ Error mengambil statistik peminjaman:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil statistik peminjaman"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func GetLoanNotifications(c *gin.Context) {
	log.Println("🚀 GetLoanNotifications dipanggil!") // Debugging pertama

	// Ambil userID dari parameter
	userIDStr := c.Param("user_id")
	log.Println("📌 User ID dari parameter:", userIDStr) // Debugging kedua

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println("❌ Error konversi User ID:", err) // Debugging ketiga
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID tidak valid"})
		return
	}

	// Ambil data peminjaman yang hampir jatuh tempo
	loans, err := repositories.GetPendingLoans(uint(userID))
	if err != nil {
		log.Println("❌ Error mengambil data peminjaman:", err) // Debugging keempat
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data peminjaman"})
		return
	}

	log.Println("✅ Data peminjaman ditemukan:", loans) // Debugging kelima

	// Kirimkan hasilnya sebagai respons JSON
	c.JSON(http.StatusOK, gin.H{"loans": loans})
}
