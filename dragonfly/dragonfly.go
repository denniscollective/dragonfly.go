package dragonfly

import (
	"os"
)

func ImageFor(jobstr string) (*os.File, error) {
	job, err := Decode(jobstr)

	if err != nil {
		return nil, err
	}

	file, err := job.Apply()

	return file, err
}
