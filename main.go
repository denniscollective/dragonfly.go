package main

import (
	"fmt"
	"github.com/denniscollective/dragonfly.go/dragonfly"
)

func main() {
	job, err := dragonfly.Decode(dragonfly.Stub)

	if err != nil {
		panic(err)
	}

	name, err := job.Apply()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("meow? %s\n", name)
}
