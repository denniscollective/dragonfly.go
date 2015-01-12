package dragonfly

import (
	//"fmt"
	"os"
)

type Job struct {
	Steps []Step
	Temp  *os.File
}

func (job *Job) Apply() (temp *os.File, err error) {

	temp, err = job.applyStep(0, nil)

	if err != nil {
		return
	}

	pipe := make(chan *os.File, 1)
	pipe <- temp
	defer close(pipe)
	temp, err = job.applyStep(1, pipe)
	return

}

func (job *Job) applyStep(i int, seed chan *os.File) (temp *os.File, err error) {
	fileChan, errChan := job.Steps[i].Apply(seed)

	select {
	case err = <-errChan:
	case temp = <-fileChan:
	}

	return temp, err
}
