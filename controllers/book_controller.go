package controllers

import (
	"log"
	"net/http"
	"perpustakaan-x-cgpt/config"
	"perpustakaan-x-cgpt/models"
	"perpustakaan-x-cgpt/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

// Tambah buku
func AddBookHandler(c *gin.Context) {
	var book models.Book

	// Bind JSON ke struct Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Panggil fungsi AddBook dari repository dengan db sebagai parameter pertama
	if err := repositories.AddBook(config.DB, &book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan buku"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Buku berhasil ditambahkan", "book": book})
}

// Update buku
func UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	if err := repositories.UpdateBook(uint(id), &book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil diperbarui", "book": book})
}

// Hapus buku
func DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := repositories.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}

// Endpoint untuk pencarian buku berdasarkan judul
func SearchBooks(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Judul buku harus diberikan"})
		return
	}

	books, err := repositories.SearchBooksByTitle(title)
	if err != nil {
		log.Println("❌ Error mencari buku:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencari buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}
