package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Initialize the todo service
	todoService := &TodoService{} // This will be replaced with a real implementation later

	// Inject the todo service into the app
	r.Use(func(c *gin.Context) {
		c.Set("todoService", todoService)
		c.Next()
	})

	// Define the API endpoints
	r.GET("/todos", getAllTodos)
	r.GET("/todos/:id", getTodo)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	// Full Disclosure: HTML generated with ChatGPT and I inserted some templating stuff into it.

	r.GET("/", func(c *gin.Context) {
		todos := []string{"Buy milk", "Walk the dog", "Do laundry"}
		c.HTML(http.StatusOK, "index.html", gin.H{"Todos": todos})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("*****************************")
	fmt.Printf("Server started in port %s\n", port)
	fmt.Println("*****************************\n")

	r.Run()
}

func getAllTodos(c *gin.Context) {
	todoService := c.MustGet("todoService").(TodoServiceInterface)
	todos, err := todoService.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	todoService := c.MustGet("todoService").(TodoServiceInterface)
	id := c.Param("id")
	todo, err := todoService.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func createTodo(c *gin.Context) {
	todoService := c.MustGet("todoService").(TodoServiceInterface)
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo, err := todoService.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTodo)
}

func updateTodo(c *gin.Context) {
	todoService := c.MustGet("todoService").(TodoServiceInterface)
	id := c.Param("id")
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTodo, err := todoService.UpdateTodo(id, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func deleteTodo(c *gin.Context) {
	todoService := c.MustGet("todoService").(TodoServiceInterface)
	id := c.Param("id")
	if err := todoService.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
