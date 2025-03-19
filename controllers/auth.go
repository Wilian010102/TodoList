package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "TodoList/config"
    "TodoList/models"
    "net/http"
    "os"
    "time"
)

type RegisterInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}


func Register(c *gin.Context) {
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    user := models.User{Username: input.Username, Password: string(hashedPassword)}
    config.DB.Create(&user)

    c.JSON(http.StatusOK, gin.H{"data": "Registrasi berhasil"})
}



func Login(c *gin.Context) {
    var input LoginInput
    var user models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Search User Based on Username
    config.DB.Where("username = ?", input.Username).First(&user)

    // Password verification
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
        return
    }

    // Create JWT Token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}


func Logout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
}
