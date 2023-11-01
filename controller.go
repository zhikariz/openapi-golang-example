package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Item represents an item in the CRUD system
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Items []Item

// @BasePath /api/v1

// @Summary Get a list of Items
// @Produce json
// @Success 200 {array} Item
// @Router /items [get]
func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, Items)
}

// @Summary Create a new item
// @Accept json
// @Produce json
// @Param item body Item true "Item object"
// @Success 201 {object} Item
// @Router /items [post]
func CreateItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Items = append(Items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

// @Summary Get an item by ID
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} Item
// @Router /items/{id} [get]
func GetItem(c *gin.Context) {
	id := c.Param("id")
	id2, _ := strconv.Atoi(id)
	for _, item := range Items {
		if id2 == item.ID {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// @Summary Update an item by ID
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param item body Item true "Item object"
// @Success 200 {object} Item
// @Router /items/{id} [put]
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id2, _ := strconv.Atoi(id)

	for i, item := range Items {
		if id2 == item.ID {
			Items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// @Summary Delete an item by ID
// @Produce json
// @Param id path int true "Item ID"
// @Success 204
// @Router /items/{id} [delete]
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	id2, _ := strconv.Atoi(id)
	for i, item := range Items {
		if id2 == item.ID {
			Items = append(Items[:i], Items[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
