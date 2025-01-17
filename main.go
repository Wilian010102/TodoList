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

    router := gin.Default()

    routes.SetupRoutes(router)


    port := os.Getenv("PORT")
    if port == "" {
        port = "8081" 
    }
    router.Run(":" + port)
}
