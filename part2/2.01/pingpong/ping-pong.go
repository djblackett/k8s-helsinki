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
		c.String(http.StatusOK, strconv.Itoa(requestCount))
		requestCount += 1
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	r.Run(port)
}
