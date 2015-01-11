package main

import (
	"fmt"
	"github.com/denniscollective/dragonfly.go/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := http.StartServer()
	if err != nil {
		fmt.Println(err)
	}
}
