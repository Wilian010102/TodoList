package config

import (
	"TodoList/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
    dbURL := fmt.Sprintf(
        "host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"),
        os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"),
    )
    db, err := gorm.Open("postgres", dbURL)
    if err != nil {
        panic("Gagal koneksi ke database!")
    }
    DB = db
    DB.AutoMigrate(&models.Checklist{}, &models.Item{}, &models.User{})
}
