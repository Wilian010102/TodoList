package routes

import (
    "github.com/gin-gonic/gin"
    "TodoList/controllers"
    "TodoList/middleware"
)

func SetupRoutes(router *gin.Engine) {
    v1 := router.Group("/api/v1")
    {
        // Authentication Endpoint
        v1.POST("/register", controllers.Register)
        v1.POST("/login", controllers.Login)
        v1.POST("/logout", middleware.Auth(), controllers.Logout)

        // Endpoint Checklist
        v1.POST("/list", middleware.Auth(), controllers.CreateChecklist)       // Membuat checklist
		v1.GET("/list", middleware.Auth(), controllers.GetChecklists)          // Menampilkan semua checklist
		v1.GET("/list/:id", middleware.Auth(), controllers.GetChecklistDetail) // Menampilkan detail checklist
		v1.DELETE("/list/:id", middleware.Auth(), controllers.DeleteChecklist) // Menghapus checklist

		//Endpoint Item
		v1.POST("/list/:id/item", middleware.Auth(), controllers.CreateItem)               // Membuat item to-do
		v1.GET("/item/:id", middleware.Auth(), controllers.GetItemDetail)                 // Menampilkan detail item
		v1.PUT("/item/:id", middleware.Auth(), controllers.UpdateItem)                    // Mengubah item
		v1.PATCH("/item/:id/status", middleware.Auth(), controllers.UpdateItemStatus)     // Mengubah status item
		v1.DELETE("/item/:id", middleware.Auth(), controllers.DeleteItem)                 // Menghapus item
    }
}
