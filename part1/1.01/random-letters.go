package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	var str = uuid.NewString()
	for true {
		fmt.Println(time.Now().String()+":", str)
		time.Sleep(5 * time.Second)
	}
}
