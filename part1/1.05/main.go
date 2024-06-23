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

	// Full Disclosure: HTML mostly generated with ChatGPT and I inserted some templating stuff into it.
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
