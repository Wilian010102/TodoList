package models

import (
	"gorm.io/gorm"
)

// Struct Checklist
type Checklist struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(100);not null"`
	Items []Item `json:"items" gorm:"constraint:OnDelete:CASCADE"` // Relasi dengan tabel Item
}

// Struct Item
type Item struct {
	gorm.Model
	ChecklistID uint   `json:"checklist_id"`                    // Foreign key ke tabel Checklist
	Name        string `json:"name" gorm:"type:varchar(100);not null"` // Nama item
	Status      bool   `json:"status" gorm:"default:false"`     // Status selesai (true/false)
}
