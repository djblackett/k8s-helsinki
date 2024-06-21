package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var requestCount = 0
	r := gin.Default()

	r.GET("/pingpong", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+strconv.Itoa(requestCount)+"\n")
		requestCount += 1
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run()
}
