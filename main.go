package main

import (
	"github.com/gin-gonic/gin"
	"uasgo_2024/database"
	"uasgo_2024/handlers"
)


func main() {
	router := gin.Default()

	//koneksi ke database
	database.InitDB()

	//Users
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	//Products
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProduct)
	router.POST("/products", handlers.CreateProduct)
	router.PUT("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)

	//Login
	router.POST("/login", handlers.LoginHandler)

	//Logout
	router.GET("/logout", handlers.LogoutHandler)

	// Jalankan server
	router.Run(":8080")
}
