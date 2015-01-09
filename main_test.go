package main_test

import (
	"fmt"
	"github.com/denniscollective/dragonfly.go"
	"github.com/denniscollective/dragonfly.go/dragonfly"
	"testing"
)

func TestDecodeDragonfly(t *testing.T) {
	job, err := dragonfly.Decode(main.Stub)

	if err != nil {
		t.Errorf("Deconde job got error %s", err)
	}

	fmt.Println(job)
	fmt.Println(len(job))

	if len(job) < 1 {
		t.Error("fail")
	}
}
