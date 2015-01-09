package main

import (
	"fmt"
	"github.com/denniscollective/dragonfly.go/dragonfly"
)

const Stub string = "W1siZmYiLCIvVXNlcnMvZGVubmlzL3dvcmtzcGFjZS96aXZpdHkvcHVibGljL2NvbnRlbnQvcGhvdG9zZXRzL29yaWdpbmFsc19hcmNoaXZlLzAwMC8wMDAwMDAvMDAwMDAwMDA3LzAwMDAwMDAwMjQtaC1vcmlnLmpwZyJdLFsicCIsInRodW1iIiwiMjB4MjAiXV0"

func main() {
	job, err := dragonfly.Decode(Stub)

	if err != nil {
		panic(err)
	}

	name, err := job.Apply()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("meow? %s\n", name)
}
