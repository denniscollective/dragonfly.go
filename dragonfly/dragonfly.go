package dragonfly

import (
	"os"
)

func ImageFor(jobstr string) (*os.File, error) {
	job, err := Decode(jobstr)

	if err != nil {
		panic(err)
	}

	file, err := job.Apply()

	return file, err
}
