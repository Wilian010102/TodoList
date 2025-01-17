package controllers

import (
	"TodoList/config"
	"TodoList/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


func CreateChecklist(c *gin.Context) {
	var checklist models.Checklist
	if err := c.ShouldBindJSON(&checklist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&checklist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, checklist)
}

// API untuk menghapus checklist
func DeleteChecklist(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Checklist{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Checklist berhasil dihapus"})
}

// API untuk menampilkan semua checklist
func GetChecklists(c *gin.Context) {
	var checklists []models.Checklist
	if err := config.DB.Preload("Items").Find(&checklists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, checklists)
}

// API untuk menampilkan detail checklist (beserta item)
func GetChecklistDetail(c *gin.Context) {
	id := c.Param("id")
	var checklist models.Checklist
	if err := config.DB.Preload("Items").First(&checklist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checklist tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, checklist)
}

// API untuk membuat item to-do di dalam checklist
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// API untuk menampilkan detail item
func GetItemDetail(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// API untuk mengubah item di dalam checklist
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item tidak ditemukan"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// API untuk mengubah status item di dalam checklist
func UpdateItemStatus(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item tidak ditemukan"})
		return
	}
	status, err := strconv.ParseBool(c.Query("status"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status harus berupa boolean"})
		return
	}
	item.Status = status
	if err := config.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// API untuk menghapus item dari checklist
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item berhasil dihapus"})
}