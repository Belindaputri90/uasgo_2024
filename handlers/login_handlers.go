package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"uasgo_2024/database"
)

// LoginRequest struct untuk menangkap data login dari body request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse struct untuk menanggapi hasil login
type LoginResponse struct {
	Message string `json:"message"`
}

// LoginHandler menangani permintaan login
func LoginHandler(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := database.GetUserByUsername(loginRequest.Username)
	if err != nil || user.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Berhasil login
	response := LoginResponse{Message: "Login successful"}
	c.JSON(http.StatusOK, gin.H{"data": response})
}


