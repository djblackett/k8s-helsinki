package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var apiUrl = os.Getenv("API_URL")

	result, err := http.Get("https://en.wikipedia.org/wiki/Special:Random")

	if err != nil {
		panic(err)
	}

	var loco = result.Request.URL

	todo := Todo{
		Text:      `Read ` + loco.String(),
		Completed: false,
	}

	fmt.Println(todo.Text)

	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status code:", resp.StatusCode)
}

type Todo struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}
