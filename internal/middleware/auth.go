package middleware

import (
	"digital-wallet/internal/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Butuh login (Token tidak ditemukan)"})
			c.Abort()
			return
		}

		// 2. Ambil tokennya saja
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// 3. Validasi tokennya
		// Karena fungsi config.ValidateToken kita sudah mengembalikan Claims,
		// kita langsung tampung hasilnya di variabel 'claims'
		claims, err := config.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid: " + err.Error()})
			c.Abort()
			return
		}

		// 4. Simpan data user ke dalam Context
		// Sekarang 'claims' sudah bisa langsung diakses
		c.Set("user_id", claims["user_id"])

		c.Next()
	}
}
