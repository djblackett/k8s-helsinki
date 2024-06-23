package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"regexp"
	"time"
)

var timestamp string

func main() {
	var str = uuid.NewString()
	time.Sleep(3 * time.Second)
	filename := "tmp/kube/timestamp.txt"
	go readTimestamp(filename)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var ping = readPingPong()
		match, _ := regexp.MatchString("\\d+", ping)
		if match {
			c.String(http.StatusOK, timestamp+" "+str+"\n"+"Ping / Pongs:"+ping)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run()
}

func readTimestamp(filename string) {
	for true {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		// Create a new scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			timestamp = scanner.Text()
		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}

		time.Sleep(5 * time.Second)
	}
}

func readPingPong() string {
	var pingpong string
	file, err := os.OpenFile("tmp/kube/pingpong.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pingpong = scanner.Text()
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return pingpong
}
