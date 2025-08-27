package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Pages  int     `json:"pages"`
	Price  float64 `json:"price"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var BooksInStore = []Book{
	{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan", Pages: 380, Price: 1200},
	{ID: "2", Title: "Clean Code", Author: "Robert C. Martin", Pages: 464, Price: 950},
}

var Categories = []Category{
	{ID: "1", Name: "Programming"},
	{ID: "2", Name: "Software Engineering"},
}

func getBooks(c *gin.Context) {
	bookID := c.Query("ID")
	if bookID != "" {
		filter := []Book{}
		for _, b := range BooksInStore {
			if fmt.Sprint(b.ID) == bookID {
				filter = append(filter, b)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, BooksInStore)
}

func getCategories(c *gin.Context) {
	cateID := c.Query("ID")
	if cateID != "" {
		filter := []Category{}
		for _, cty := range Categories {
			if fmt.Sprint(cty.ID) == cateID {
				filter = append(filter, cty)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, Categories)
}

func main() {
	r := gin.Default()

	r.GET("/StoreInfo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Store":  "Mai Bookstore",
			"Slogan": "อ่านเยอะ ความรู้เยอะ",
		})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/book", getBooks)
		api.GET("/category", getCategories)
	}

	r.Run(":8080")
}
