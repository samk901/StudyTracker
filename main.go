package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", getUsers)
		v1.POST("/users", createUser)
	}

	router.Run()
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": []User{
			{ID: "1", Name: "John"},
		},
	})
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add User logic here
	fmt.Println("Hello World!")

	c.JSON(http.StatusCreated, newUser)
}
