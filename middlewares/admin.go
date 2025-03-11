package middlewares

import (
	"net/http"
	"perpustakaan-x-cgpt/repositories"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Claims adalah struktur untuk menyimpan data dari token JWT
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Middleware untuk memastikan hanya admin yang bisa mengakses endpoint
func AdminOnly(c *gin.Context) {
	email := c.Request.Header.Get("Admin-Email")
	if email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Akses ditolak, hanya admin yang bisa"})
		c.Abort()
		return
	}

	user, err := repositories.GetUserByEmail(email)
	if err != nil || user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya admin yang bisa mengakses"})
		c.Abort()
		return
	}

	c.Next()
}
