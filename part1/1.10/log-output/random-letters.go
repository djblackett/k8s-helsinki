package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
)

func main() {
	var str = uuid.NewString()
	time.Sleep(3 * time.Second)
	filename := "files/timestamp.txt"

	// Open the file

	for true {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Create a new scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text(), str)
		}

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}

		time.Sleep(5 * time.Second)
	}
}
