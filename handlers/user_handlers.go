package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"uasgo_2024/database"
	"uasgo_2024/models"
)

// UserGinH is a type to replace gin.H for Swagger accessibility
type UserGinH map[string]interface{}

// @Summary Get list of users
// @Description Get list of users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} UserGinH
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	// Return only username and password
	var response []UserGinH
	for _, user := range users {
		response = append(response, UserGinH{"username": user.Username, "password": user.Password})
	}

	c.JSON(http.StatusOK, UserGinH{"data": response})
}

// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserGinH
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	database.DB.First(&user, id)

	// Return only username and password
	response := UserGinH{"username": user.Username, "password": user.Password}
	c.JSON(http.StatusOK, UserGinH{"data": response})
}

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param input body models.User true "User input"
// @Success 200 {object} UserGinH
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, UserGinH{"error": err.Error()})
		return
	}

	database.DB.Create(&userInput)
	response := UserGinH{"username": userInput.Username, "password": userInput.Password}
	c.JSON(http.StatusOK, UserGinH{"data": response})
}

// @Summary Update a user by ID
// @Description Update a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body models.User true "User input"
// @Success 200 {object} UserGinH
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, UserGinH{"error": err.Error()})
		return
	}

	var user models.User
	database.DB.First(&user, id)
	user.Username = userInput.Username
	user.Password = userInput.Password
	database.DB.Save(&user)

	response := UserGinH{"username": user.Username, "password": user.Password}
	c.JSON(http.StatusOK, UserGinH{"data": response})
}

// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} UserGinH
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.User{}, id)

	c.JSON(http.StatusOK, UserGinH{"data": "User deleted"})
}
