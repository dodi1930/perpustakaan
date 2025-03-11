package repositories

import (
	"perpustakaan-x-cgpt/config"
	"perpustakaan-x-cgpt/models"

	"gorm.io/gorm"
)

// Tambah buku
func AddBook(db *gorm.DB, book *models.Book) error {
	return db.Create(book).Error
}

// Update buku
func UpdateBook(id uint, book *models.Book) error {
	return config.DB.Model(&models.Book{}).Where("id = ?", id).Updates(book).Error
}

// Hapus buku
func DeleteBook(id uint) error {
	return config.DB.Delete(&models.Book{}, id).Error
}

// Cari buku berdasarkan judul (menggunakan LIKE agar bisa mencari sebagian kata)
func SearchBooksByTitle(title string) ([]models.Book, error) {
	var books []models.Book
	result := config.DB.Where("title ILIKE ?", "%"+title+"%").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
