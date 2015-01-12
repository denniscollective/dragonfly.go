package dragonfly

import (
	"os"
)

func ImageFor(jobstr string) (file *os.File, err error) {
	job, err := Decode(jobstr)

	if err != nil {
		return nil, err
	}

	file, err = job.Apply()

	return
}
