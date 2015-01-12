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
	seed := make(chan *os.File)
	close(seed)

	seed, errChan := job.Steps[0].Apply(seed)
	seed2, errChan2 := job.Steps[1].Apply(seed)

	select {
	//case err = <-errChan:

	case err = <-errChan:
		close(seed2)
		close(errChan2)
	case err = <-errChan2:
	case temp = <-seed:

	}
	return

}
