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

	// Pick user_id from validated token JWT
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Bind JSON to struct list
	if err := c.ShouldBindJSON(&checklist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set UserID from user login
	checklist.UserID = userID.(uint)

	//Save list to database
	if err := config.DB.Create(&checklist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, checklist)
}

//API for update list
func UpdateChecklist(c *gin.Context) {
	id := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var checklist models.Checklist

	// Validate list is exist and owned by user login
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&checklist).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found or not yours"})
		return
	}

	// Bind JSON to struct checklist
	if err := c.ShouldBindJSON(&checklist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save change to database
	if err := config.DB.Save(&checklist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update checklist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checklist berhasil diperbarui", "checklist": checklist})
}


// API for delete list
func DeleteChecklist(c *gin.Context) {
	id := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var checklist models.Checklist
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&checklist).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found or not yours"})
		return
	}

	if err := config.DB.Delete(&checklist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete checklist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List successfully deleted"})
}

// API for show all list
func GetChecklists(c *gin.Context) {
	var checklists []models.Checklist

	// Pick user_id from JWT token
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Pick list that created by user login
	if err := config.DB.Preload("Items").Where("user_id = ?", userID).Find(&checklists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to take data"})
		return
	}

	c.JSON(http.StatusOK, checklists)
}

// API for show list detail
func GetChecklistDetail(c *gin.Context) {
	id := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var checklist models.Checklist
	if err := config.DB.Preload("Items").Where("id = ? AND user_id = ?", id, userID).First(&checklist).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found or not yours"})
		return
	}

	c.JSON(http.StatusOK, checklist)
}

// API for create item (only for owned list)
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Validate list is exist and owned by user login
	var checklist models.Checklist
	if err := config.DB.Where("id = ? AND user_id = ?", item.ChecklistID, userID).First(&checklist).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "List not found or not yours"})
		return
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// API for show item detail
func GetItemDetail(c *gin.Context) {
	id := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var item models.Item
	if err := config.DB.Joins("JOIN checklists ON checklists.id = items.checklist_id").
		Where("items.id = ? AND checklists.user_id = ?", id, userID).
		First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found or not yours"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// API for update item (only for owned item)
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var item models.Item
	if err := config.DB.Joins("JOIN checklists ON checklists.id = items.checklist_id").
		Where("items.id = ? AND checklists.user_id = ?", id, userID).
		First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found or not yours"})
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

// API for update item status (only for owned item)
func UpdateItemStatus(c *gin.Context) {
	id := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var item models.Item
	if err := config.DB.Joins("JOIN checklists ON checklists.id = items.checklist_id").
		Where("items.id = ? AND checklists.user_id = ?", id, userID).
		First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found or not yours"})
		return
	}

	status, err := strconv.ParseBool(c.Query("status"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status must be true or false"})
		return
	}

	item.Status = status
	if err := config.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// API for delete item (only for owned item)
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var item models.Item
	if err := config.DB.Joins("JOIN checklists ON checklists.id = items.checklist_id").
		Where("items.id = ? AND checklists.user_id = ?", id, userID).
		First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found or not yours"})
		return
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item successfully deleted"})
}