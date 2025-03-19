package middleware

import (
    "net/http"
    "strings"
    "TodoList/utils"
    "github.com/gin-gonic/gin"
    "log"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            log.Println("Error: Header Authorization tidak ditemukan")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Header Authorization tidak ditemukan"})
            c.Abort()
            return
        }

        if !strings.HasPrefix(tokenString, "Bearer ") {
            log.Println("Error: Format token tidak valid. Harus diawali dengan 'Bearer '")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token tidak valid. Harus diawali dengan 'Bearer '"})
            c.Abort()
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        log.Println("Token Diterima:", tokenString)

        // Verifikasi token JWT
        claims, err := utils.VerifyToken(tokenString)
        if err != nil {
            log.Println("Error: Token tidak valid atau sudah kadaluarsa -", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau sudah kadaluarsa"})
            c.Abort()
            return
        }

        log.Println("Autentikasi berhasil, user_id:", claims.UserID)
        c.Set("user_id", claims.UserID)
        c.Next()
    }
}