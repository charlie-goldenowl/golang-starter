package controllers

import (
	"github.com/charlie-goldenowl/golangstarter1/initializers"
	"github.com/charlie-goldenowl/golangstarter1/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	db := initializers.ConnectDB(&config)
	books := []models.Book{}
	dberr := db.Debug().Model(&models.Book{}).Limit(100).Find(&books).Error
	if dberr != nil {
		c.JSON(http.StatusBadGateway, gin.H{"data": dberr})
		return
	}
	log.Println()
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create new books
func CreateBook(c *gin.Context) {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	db := initializers.ConnectDB(&config)
	// Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Book
	book := models.Book{Title: input.Author, Author: input.Author}
	db.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
