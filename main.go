package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	_ "study_tracker/docs" // replace with your actual module name

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *sql.DB

func initDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	// Get environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to database!")
	return db, nil
}

// Database configs
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

// Database queries
const (
	createUserQuery = `
		INSERT INTO Users (name) 
		VALUES ($1) 
		RETURNING id`

	getUsersQuery = `
		SELECT id, name 
		FROM Users`

	createQuestionQuery = `
		INSERT INTO Questions (name, pattern, difficulty, last_completed_time, next_due_time, notes) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id`

	getQuestionsQuery = `
		SELECT id, name, pattern, difficulty, last_completed_time, next_due_time, notes 
		FROM Questions`
)

func main() {

	var err error
	db, err = initDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

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
	rows, err := db.Query(getUsersQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
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

	err := db.QueryRow(createUserQuery, newUser.Name).Scan(&newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

// @Summary     Get questions
// @Description Get list of questions
// @Tags        questions
// @Produce     json
// @Success     200 {object} map[string][]LeetCodeQuestions "Questions list response"
// @Router      /questions [get]
func getQuestions(c *gin.Context) {
	rows, err := db.Query(getQuestionsQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	defer rows.Close()

	var questions []LeetCodeQuestions
	for rows.Next() {
		var q LeetCodeQuestions
		if err := rows.Scan(&q.ID, &q.Name, &q.Pattern, &q.Difficulty, &q.LastCompletedTime, &q.NextDueTime, &q.Notes); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		questions = append(questions, q)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"questions": questions})
}

// @Summary     Create question
// @Description Create a new question
// @Tags        questions
// @Accept      json
// @Produce     json
// @Param       question body LeetCodeQuestions true "Question object"
// @Success     201 {object} LeetCodeQuestions "Question created"
// @Failure     400 {object} ErrorResponse "Bad request"
// @Router      /questions [post]
func createQuestion(c *gin.Context) {
	var newQuestion LeetCodeQuestions
	if err := c.BindJSON(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.QueryRow(createQuestionQuery, newQuestion.Name, newQuestion.Pattern, newQuestion.Difficulty, newQuestion.LastCompletedTime, newQuestion.NextDueTime, newQuestion.Notes).Scan(&newQuestion.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newQuestion)
}
