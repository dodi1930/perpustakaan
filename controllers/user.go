package controllers

import (
	"fmt"
	"net/http"
	"perpustakaan-x-cgpt/config"
	"perpustakaan-x-cgpt/models"
	"perpustakaan-x-cgpt/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// AddUser handler
func AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah email sudah terdaftar
	existingUser, _ := repositories.GetUserByEmail(user.Email)
	if existingUser.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email sudah terdaftar"})
		return
	}

	// Hash password sebelum disimpan
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal meng-hash password"})
		return
	}

	// Simpan user ke database
	err := repositories.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan user", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil ditambahkan!"})
}

// DeleteUser handler
func DeleteUser(c *gin.Context) {
	// Ambil ID user yang akan dihapus dari URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Hapus user dari database
	err = repositories.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus!"})
}

// EditUser handler
func EditUser(c *gin.Context) {
	// Ambil ID user yang akan diedit dari URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Ambil data baru dari request
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update data user
	err = repositories.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil diperbarui!"})
}

// ImportUsers meng-handle upload dan proses import user dari Excel
func ImportUsers(c *gin.Context) {
	// Ambil file dari request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengambil file"})
		return
	}

	// Simpan sementara
	filePath := "./temp.xlsx"
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	// Buka file Excel
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuka file Excel"})
		return
	}
	defer f.Close()

	// Baca sheet pertama
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca sheet"})
		return
	}

	// Iterasi setiap baris, mulai dari baris ke-2 (header dilewati)
	for i, row := range rows {
		if i == 0 {
			continue // Lewati header
		}

		if len(row) < 3 {
			continue // Lewati baris yang tidak lengkap
		}

		user := models.User{
			Name:  row[0],
			Email: row[1],
			Role:  row[3], // Role
		}

		// Hash password sebelum disimpan
		user.Password = row[2]
		if err := user.HashPassword(); err != nil {
			fmt.Println("Gagal meng-hash password:", err)
			continue
		}

		// Simpan ke database
		err := repositories.CreateUser(&user)
		if err != nil {
			fmt.Println("Gagal menambahkan user:", err)
			continue
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Import berhasil"})
}

// GetUsers mengambil semua pengguna
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pengguna"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
