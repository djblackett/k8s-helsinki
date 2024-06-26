package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var uniqueId = 3

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"}, // Adjust to your React app's address
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	todos := []Todo{
		{Id: 0, Text: "Welcome to your new todo list", Completed: false},
		{
			Id:        1,
			Text:      "Tap the sun to switch to light mode",
			Completed: false,
		},
		{
			Id:        2,
			Text:      "Tap the circles to mark items Completed",
			Completed: false,
		}}

	r.GET("/todos", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		var newTodo Todo

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&newTodo); err != nil {
			return
		}

		newTodo.Id = uniqueId
		uniqueId += 1

		// Add the new album to the slice.
		todos = append(todos, newTodo)
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
