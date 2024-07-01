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
		requestCount += 1
		c.String(http.StatusOK, strconv.Itoa(requestCount))

	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, strconv.Itoa(requestCount))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	r.Run(port)
}
