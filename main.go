package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"uasgo_2024/database"
	"uasgo_2024/handlers"
	_ "uasgo_2024/docs" // Import package docs for swag init
)

// @title Your API Title
// @version 1.0
// @description Your API description. This can be multiline.
// @host localhost:8080
// @BasePath /v1
func main() {
	router := gin.Default()

	// Koneksi ke database
	database.InitDB()

	// Endpoint Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoint Users
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	// Endpoint Products
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProduct)
	router.POST("/products", handlers.CreateProduct)
	router.PUT("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)

	// Endpoint Login
	router.POST("/login", handlers.LoginHandler)

	// Endpoint Logout
	router.GET("/logout", handlers.LogoutHandler)

	// Jalankan server
	router.Run(":8080")
}
