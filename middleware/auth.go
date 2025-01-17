package middleware

import (
    "net/http"
    "strings"
    "TodoList/utils"
    "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Ambil header Authorization
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Header Authorization tidak ditemukan"})
            c.Abort()
            return
        }

        // Validasi format token (harus diawali dengan "Bearer ")
        if !strings.HasPrefix(tokenString, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token tidak valid. Harus diawali dengan 'Bearer '"})
            c.Abort()
            return
        }

        // Hilangkan prefix "Bearer " untuk mendapatkan token sebenarnya
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Verifikasi token menggunakan utils
        claims, err := utils.VerifyToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau sudah kadaluarsa"})
            c.Abort()
            return
        }

        // Simpan informasi pengguna ke dalam context (opsional)
        c.Set("user_id", claims.UserID)

        // Lanjutkan ke handler berikutnya
        c.Next()
    }
}
