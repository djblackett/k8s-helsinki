package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var str = uuid.NewString()

	r := gin.Default()
	url := os.Getenv("PING_PONG_SERVICE")

	if url == "" {
		log.Fatal("PING_PONG_SERVICE environment variable not set")
	}

	r.GET("/", func(c *gin.Context) {
		var timestamp = time.Now().Format(time.RFC3339)
		resp, err := http.Get(url + "/pingpong")

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		c.String(http.StatusOK, timestamp+" "+str+"\n"+"Ping / Pongs: "+string(bodyBytes))

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run()
}
