package main

import (
	"go-bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: "1", Title: "Go Programming", Author: "John Doe", Price: 39.99},
	{ID: "2", Title: "Mastering Gin", Author: "Jane Smith", Price: 29.99},
}

// Get all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Get a book by ID
func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// Add a new book
func addBook(c *gin.Context) {
	var newBook models.Book

	// Bind JSON request body to newBook
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// Update a book by ID
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book

	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// Delete a book by ID
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	router := gin.Default()

	// Define API routes
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", addBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	// Run the server on port 8080
	router.Run("localhost:8081")
}
