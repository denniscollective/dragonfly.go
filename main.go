package main

import (
	"fmt"
	"github.com/denniscollective/dragonfly.go/http"
)

func main() {
	err := http.StartServer()
	if err != nil {
		fmt.Println(err)
	}
}
