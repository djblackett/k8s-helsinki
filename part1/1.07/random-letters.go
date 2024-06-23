package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"
)

var timestamp string

func main() {
	var str = uuid.NewString()
	go stuff(str)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"text": str, "timestamp": timestamp})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run()
}

func stuff(str string) {
	for true {
		timestamp = time.Now().String()
		fmt.Println(timestamp+":", str)
		time.Sleep(5 * time.Second)
	}
}
