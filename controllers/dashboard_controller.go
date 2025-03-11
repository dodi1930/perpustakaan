package controllers

import (
	"log"
	"net/http"
	"perpustakaan-x-cgpt/models"

	"github.com/gin-gonic/gin"
)

func DashboardHandler(c *gin.Context) {
	var totalBuku, totalPengguna, peminjamanAktif int64

	// Debugging: Cek apakah database terhubung
	if models.DB == nil {
		log.Println("Koneksi database tidak tersedia!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database tidak terhubung"})
		return
	}

	// Cek error query database
	if err := models.DB.Table("books").Count(&totalBuku).Error; err != nil {
		log.Println("Error menghitung jumlah buku:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung jumlah buku"})
		return
	}

	if err := models.DB.Table("users").Count(&totalPengguna).Error; err != nil {
		log.Println("Error menghitung jumlah pengguna:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung jumlah pengguna"})
		return
	}

	if err := models.DB.Table("loans").Where("status = ?", "dipinjam").Count(&peminjamanAktif).Error; err != nil {
		log.Println("Error menghitung jumlah peminjaman:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung jumlah peminjaman"})
		return
	}

	// Debugging: Pastikan template ada
	if _, err := c.Get("dashboard.html"); err {
		log.Println("Template dashboard.html tidak ditemukan!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Template dashboard.html tidak ditemukan"})
		return
	}

	// Jika semua berhasil, render template
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"totalBuku":       totalBuku,
		"totalPengguna":   totalPengguna,
		"peminjamanAktif": peminjamanAktif,
	})
}
