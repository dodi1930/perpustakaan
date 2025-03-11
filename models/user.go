package models

import (
	//"net/http"

	//"perpustakaan-x-cgpt/utils"
	"time"

	//"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"` // Jangan kirim password ke JSON response
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// Fungsi untuk Hash Password sebelum disimpan
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword) // Simpan hasil hash ke struct
	return nil
}
