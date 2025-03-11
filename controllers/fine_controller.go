package controllers

import (
	"net/http"
	"strconv"
	"time"

	"perpustakaan-x-cgpt/models"
	"perpustakaan-x-cgpt/repositories"

	"github.com/gin-gonic/gin"
)

// Handler untuk mendapatkan denda per hari
func GetFinePerDay(c *gin.Context) {
	fine, err := repositories.GetFinePerDay()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data denda", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"fine_per_day": fine})
}

// Handler untuk memperbarui denda per hari
func UpdateFinePerDay(c *gin.Context) {
	// Ambil nilai baru dari parameter request
	newFineStr := c.PostForm("fine_per_day")
	newFine, err := strconv.ParseFloat(newFineStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format nilai denda tidak valid"})
		return
	}

	// Update denda di database
	err = repositories.UpdateFinePerDay(newFine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui denda", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Denda berhasil diperbarui", "new_fine_per_day": newFine})
}

// GetUserFine menghitung total denda seorang pengguna berdasarkan keterlambatan pengembalian buku
func GetUserFine(userID int) (float64, error) {
	var totalFine float64

	// Ambil daftar peminjaman yang terlambat dikembalikan
	var borrowedBooks []models.BorrowedBook
	if err := models.DB.Where("user_id = ? AND return_date > due_date", userID).Find(&borrowedBooks).Error; err != nil {
		return 0, err
	}

	// Ambil nilai denda per hari dari repository, bukan dari handler
	finePerDay, err := repositories.GetFinePerDay()
	if err != nil {
		return 0, err
	}

	// Hitung total denda berdasarkan jumlah hari keterlambatan
	for _, book := range borrowedBooks {
		lateDays := int(time.Since(book.DueDate).Hours() / 24) // Hitung jumlah hari keterlambatan
		if lateDays > 0 {
			totalFine += finePerDay * float64(lateDays)
		}
	}

	return totalFine, nil
}

// Handler untuk mendapatkan denda seorang pengguna berdasarkan ID
func GetUserFineHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
		return
	}

	totalFine, err := GetUserFine(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung denda", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "total_fine": totalFine})
}
