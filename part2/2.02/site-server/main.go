package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	_, err := readTimestamp()
	if err != nil {
		writeTimestamp()
		getImage()
	}

	go startTimestampWatcher()

	r := gin.Default()

	r.Static("/todo-react", "./build") // Serve static files from React's build directory
	r.NoRoute(func(c *gin.Context) {
		c.File("./build/index.html") // Serve the React app for any undefined routes
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	fmt.Println("*****************************")
	fmt.Printf("Server started in port %s\n", port)
	fmt.Println("*****************************\n")

	r.Run(port)
}

func getImage() {
	resp, err := http.Get("https://picsum.photos/1200")
	if err != nil {
		log.Fatalln(err)
	}

	filename := "tmp/kube/img.jpg"

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Write the timestamp to the file
	if _, err := io.Copy(file, resp.Body); err != nil {
		fmt.Println("Error writing to file:", err)
	}

	fmt.Println("Image updated")
}

func readTimestamp() (string, error) {
	var filename = "tmp/kube/timestamp.txt"
	var timestamp string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
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

	return timestamp, nil
}

func writeTimestamp() {
	var filename = "tmp/kube/timestamp.txt"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	timestamp := time.Now().Format(time.RFC3339)

	file.Truncate(0)
	file.Seek(0, 0)
	// Write the timestamp to the file
	if _, err := file.WriteString(timestamp + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Timestamp written to file:", timestamp)
}

func checkTimestamp() bool {
	timestamp, err1 := readTimestamp()
	if err1 != nil {
		fmt.Println("error reading timestamp")
	}
	t, err2 := time.Parse(time.RFC3339, timestamp)
	if err2 != nil {
		fmt.Println("error parsing timestamp")
	}
	return time.Since(t) > time.Hour
}

func startTimestampWatcher() {
	fmt.Println("Starting timestamp observer")
	for true {
		if checkTimestamp() {
			writeTimestamp()
			getImage()
		}
		time.Sleep(1 * time.Hour)
	}
}
