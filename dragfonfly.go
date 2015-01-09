package main

import (
	"encoding/base64"
	"fmt"
)

const stub string = "W1siZmYiLCIvVXNlcnMvZGVubmlzL3dvcmtzcGFjZS96aXZpdHkvcHVibGljL2NvbnRlbnQvcGhvdG9zZXRzL29yaWdpbmFsc19hcmNoaXZlLzAwMC8wMDAwMDAvMDAwMDAwMDA3LzAwMDAwMDAwMjQtaC1vcmlnLmpwZyJdLFsicCIsInRodW1iIiwiMjB4MjAiXV0"

func main() {

	json, err := decodeJob(stub)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("meow? %s\n", json)
}

func decodeJob(jobStr string) (*[]byte, error) {
	job := jobStr + "=\n"
	data, err := base64.StdEncoding.DecodeString(job) //dragonfly trims a trailing =\n from the jobs
	if err != nil {
		return nil, err
	}

	return &data, err
}
