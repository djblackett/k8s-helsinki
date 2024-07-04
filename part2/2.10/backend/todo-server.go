package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		logger.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

func main() {

	var host = os.Getenv("HOST")
	var password = os.Getenv("PASSWORD") // switch to encrypted secret later
	postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic("Invalid port number")
	}
	var dbname = os.Getenv("DB_NAME")
	var user = os.Getenv("USER")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, postgresPort, user, password, dbname)

	fmt.Println(psqlInfo)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		return
	}

	fmt.Println("Successfully connected!")

	r := gin.Default()
	r.Use(LoggerMiddleware())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		result := db.Find(&todos)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		var newTodo Todo

		// Call BindJSON to bind the received JSON
		if err := c.BindJSON(&newTodo); err != nil {
			return
		}

		db.Create(&newTodo)
		fmt.Println(newTodo)
		c.IndentedJSON(http.StatusCreated, newTodo)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	r.Run(port)

}

type Todo struct {
	gorm.Model
	Id        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

func (todo Todo) completeTodo(todos []Todo, Id int) {
	for i, todo := range todos {
		if todo.Id == Id {
			todos[i].Completed = true
			fmt.Println("Todo item Completed:", todos[i])
			return
		}
	}
	fmt.Println("Todo item with Id", Id, "not found")
}
