package models

import "gorm.io/gorm"

// User represents a user model
// @Description User model
type User struct {
    gorm.Model
    Username string `json:"username" gorm:"unique"`
    Password string `json:"password"`
}
