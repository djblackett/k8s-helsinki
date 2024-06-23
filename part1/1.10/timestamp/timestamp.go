package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filename := "files/timestamp.txt"

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	for true {
		timestamp := time.Now().Format(time.RFC3339)

		file.Truncate(0)
		file.Seek(0, 0)
		// Write the timestamp to the file
		if _, err := file.WriteString(timestamp + "\n"); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		time.Sleep(5 * time.Second)
	}
	defer file.Close()
}
