package repositories

import (
	"fmt"
	"perpustakaan-x-cgpt/config"
	"perpustakaan-x-cgpt/models"
	"time"
)

// CreateUser untuk menambahkan user baru
func CreateUser(user *models.User) error {
	if err := user.HashPassword(); err != nil {
		fmt.Println("Error hashing password:", err)
		return err
	}

	// Gunakan GORM untuk insert data
	user.CreatedAt = time.Now()
	err := config.DB.Create(&user).Error
	if err != nil {
		fmt.Println("Error inserting user into database:", err)
	}
	return err
}

// GetUserByEmail mencari user berdasarkan email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// UpdateUser memperbarui data user berdasarkan ID, termasuk password jika diubah
func UpdateUser(id int, user models.User) error {
	// Ambil data user lama dari database
	var existingUser models.User
	if err := config.DB.First(&existingUser, id).Error; err != nil {
		return err
	}

	// Jika password kosong, gunakan password lama
	if user.Password == "" {
		user.Password = existingUser.Password
	} else {
		// Hash password baru sebelum disimpan
		if err := user.HashPassword(); err != nil {
			return err
		}
	}

	// Update user di database
	return config.DB.Model(&existingUser).Updates(user).Error
}

// DeleteUser menghapus user berdasarkan ID
func DeleteUser(id int) error {
	return config.DB.Delete(&models.User{}, id).Error
}

// BatchCreateUsers untuk import user dalam jumlah banyak
func BatchCreateUsers(users []models.User) error {
	return config.DB.Create(&users).Error
}
