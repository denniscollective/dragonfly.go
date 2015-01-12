package dragonfly

import (
	//"fmt"
	"os"
)

type Job struct {
	Steps []*Step
	Temp  *os.File
}

func (job *Job) Apply() (temp *os.File, err error) {
	fileChan := make(chan *os.File, 1)
	errChan := make(chan error, 1)
	defer close(fileChan)
	defer close(errChan)

	go job.Steps[0].Fetch(fileChan, errChan)

	select {
	case err = <-errChan:
		return nil, err
	case temp = <-fileChan:
	}

	go job.Steps[1].Process(temp, fileChan, errChan)
	select {
	case err = <-errChan:
		return nil, err
	case temp = <-fileChan:
	}

	return
}
