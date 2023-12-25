package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"bookstore/pkg/models"
	"github.com/gin-gonic/gin"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	c.IndentedJSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	fmt.Println("bookId = " + bookId)
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	c.IndentedJSON(http.StatusOK, bookDetails)
}

func CreateBook(c *gin.Context) {
	newBook := &models.Book{}

	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println(err)
		return
	}

	result := newBook.CreateBook()
	c.IndentedJSON(http.StatusCreated, result)
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	result := models.DeleteBook(ID)
	c.IndentedJSON(http.StatusOK, result)
}

func UpdateBook(c *gin.Context) {
	var updateBook = &models.Book{}

	if err := c.BindJSON(&updateBook); err != nil {
		fmt.Println(err)
		return
	}

	bookId := c.Param("bookId")
	fmt.Println("bookId = " + bookId)
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	c.IndentedJSON(http.StatusAccepted, bookDetails)
}
