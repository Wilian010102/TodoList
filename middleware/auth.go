package middleware

import (
    "net/http"
    "strings"
    "TodoList/utils"
    "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
       
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Header Authorization tidak ditemukan"})
            c.Abort()
            return
        }

        
        if !strings.HasPrefix(tokenString, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token tidak valid. Harus diawali dengan 'Bearer '"})
            c.Abort()
            return
        }

       
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")


        claims, err := utils.VerifyToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau sudah kadaluarsa"})
            c.Abort()
            return
        }


        c.Set("user_id", claims.UserID)


        c.Next()
    }
}
