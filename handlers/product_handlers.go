package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"uasgo_2024/database"
	"uasgo_2024/models"
)

// ProductGinH is a type to replace gin.H for Swagger accessibility
type ProductGinH map[string]interface{}

// @Summary Get list of products
// @Description Get list of products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} ProductGinH
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)

	c.JSON(http.StatusOK, ProductGinH{"data": products})
}

// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} ProductGinH
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	database.DB.First(&product, id)

	c.JSON(http.StatusOK, ProductGinH{"data": product})
}

// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param input body models.Product true "Product input"
// @Success 200 {object} ProductGinH
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var productInput models.Product
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, ProductGinH{"error": err.Error()})
		return
	}

	database.DB.Create(&productInput)
	c.JSON(http.StatusOK, ProductGinH{"data": productInput})
}

// @Summary Update a product by ID
// @Description Update a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param input body models.Product true "Product input"
// @Success 200 {object} ProductGinH
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var productInput models.Product
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, ProductGinH{"error": err.Error()})
		return
	}

	var product models.Product
	database.DB.First(&product, id)
	product.Name = productInput.Name
	product.Price = productInput.Price
	database.DB.Save(&product)

	c.JSON(http.StatusOK, ProductGinH{"data": product})
}

// @Summary Delete a product by ID
// @Description Delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} ProductGinH
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.Product{}, id)

	c.JSON(http.StatusOK, ProductGinH{"data": "Product deleted"})
}
