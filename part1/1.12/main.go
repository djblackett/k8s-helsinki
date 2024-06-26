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
	r.Static("/static", "./tmp/kube")
	r.LoadHTMLGlob("templates/*")

	// Full Disclosure: HTML mostly generated with ChatGPT and I inserted some templating stuff into it.
	// I made a React app for later versions of this project (100% written by me)
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		timestamp = scanner.Text()
	}

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
