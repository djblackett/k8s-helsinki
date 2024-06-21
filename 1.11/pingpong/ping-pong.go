package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var requestCount int64
	r := gin.Default()

	filename := "tmp/kube/pingpong.txt"

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var number = scanner.Text()
		match, _ := regexp.MatchString("\\d+", number)
		if match {
			requestCount, _ = strconv.ParseInt(number, 10, 64)
		} else {
			requestCount = 0
		}

	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	r.GET("/pingpong", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+strconv.Itoa(int(requestCount))+"\n")
		requestCount += 1
		file.Truncate(0)
		file.Seek(0, 0)
		// Write the timestamp to the file
		if _, err := file.WriteString(strconv.Itoa(int(requestCount)) + "\n"); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run()
}
