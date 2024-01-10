package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//logout
func LogoutHandler(c *gin.Context) {
	
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
