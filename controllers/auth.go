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

// Registergodoc
// @Summary Registrasi User
// @Description Registrasi User supaya terdaftar di database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Users body models.User true "Data User"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register [post]
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


// Logingodoc
// @Summary Login User
// @Description Login user sesuai username dan password di database
// @Tags User
// @Accept  json
// @Produce  json
// @Param Users body models.User true "Data User"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /login [post]
// User Registration Function
func Login(c *gin.Context) {
    var input LoginInput
    var user models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Search User Based on Usernama
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

// Logoutgodoc
// @Summary Logout User
// @Description Logout User
// @Tags User
// @Accept  json
// @Produce  json
// @Param Users body models.User true "Data User"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /logout [post]
func Logout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
}
