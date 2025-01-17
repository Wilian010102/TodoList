package main

import (
    "github.com/gin-gonic/gin" 
    "TodoList/config"
    "TodoList/routes"
    "os"
	"github.com/joho/godotenv"
)

func main() {
    // Load .env
    err := godotenv.Load()
    if err != nil {
        panic("Gagal memuat file .env")
    }
    config.ConnectDB()

    // Inisialisasi router Gin
    router := gin.Default()

    // Inisialisasi routing lainnya
    routes.SetupRoutes(router)

    // Jalankan server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8081" // default port jika tidak ada di .env
    }
    router.Run(":" + port)
}
