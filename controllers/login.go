package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handler
func UserLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil!"})
}
