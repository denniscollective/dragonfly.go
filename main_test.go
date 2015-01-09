package main_test

import (
	"github.com/denniscollective/dragonfly.go"
	"github.com/denniscollective/dragonfly.go/dragonfly"
	"testing"
)

func TestDecodeDragonfly(t *testing.T) {
	job, err := dragonfly.Decode(main.Stub)

	if err != nil {
		t.Errorf("Deconde job got error %s", err)
	}

	if len(job) != 2 {
		t.Error("job should have two steps")
	}

	if job[0].Command != "ff" {
		t.Error("the first test of the stub is supposed to be fetch File")
	}

	if args := job[1].Args; args[0] != "thumb" && args[1] != "20x20" {
		t.Error("second step should be a resize to thumbnail 20x20 job")
	}
}

func TestDecodeFailse(t *testing.T) {
	job, err := dragonfly.Decode("this is y i'm hawt")
	if err == nil {
		t.Error("Decode errors aren't propagating")
	}

	if job != nil {
		t.Error("Decode should return nil when it has an error")
	}
}
