package main

import (
	"fmt"
	"net/http"
	"time"

	_ "study_tracker/docs" // replace with your actual module name

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           LeetCode Questions API
// @version         1.0
// @description     API for managing LeetCode questions and users
// @host            localhost:8080
// @BasePath        /api/v1

// Define response models for Swagger
type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Difficulty string

const (
	Easy   Difficulty = "Easy"
	Medium Difficulty = "Medium"
	Hard   Difficulty = "Hard"
)

type LeetCodeQuestions struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	Pattern           string     `json:"pattern"`
	Difficulty        Difficulty `json:"difficulty"`
	LastCompletedTime time.Time  `json:"lastCompletedTime"`
	NextDueTime       time.Time  `json:"nextDueTime"`
	Notes             string     `json:"notes"`
}

func main() {
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", getUsers)
		v1.POST("/users", createUser)
		v1.GET("/questions", getQuestions)
		v1.POST("/questions", createQuestion)
	}

	router.Run()
}

// @Summary     Get users
// @Description Get list of users
// @Tags        users
// @Produce     json
// @Success     200 {object} map[string][]User "User list response"
// @Router      /users [get]
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": []User{
			{ID: "1", Name: "John"},
		},
	})
}

// @Summary     Create user
// @Description Create a new user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user body User true "User object"
// @Success     201 {object} User "User created"
// @Failure     400 {object} ErrorResponse "Bad request"
// @Router      /users [post]
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

func getQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"questions": []LeetCodeQuestions{
			{ID: "1", Name: "Two Sum", Pattern: "Array", Difficulty: Easy, LastCompletedTime: time.Now(), NextDueTime: time.Now(), Notes: "Notes"},
		},
	})
}

func createQuestion(c *gin.Context) {
	var newQuestion LeetCodeQuestions
	if err := c.BindJSON(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add Question logic here
	fmt.Println("Created a new question!")

	c.JSON(http.StatusCreated, newQuestion)
}
