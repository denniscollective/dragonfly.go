package dragonfly

import (
	//"fmt"
	"os"
)

type Job struct {
	Steps []Step
	Temp  *os.File
}

func (job Job) Apply() (string, error) {
	fileChan := make(chan *os.File)
	errChan := make(chan error)
	go job.Steps[0].Fetch(fileChan, errChan)

	var (
		temp *os.File
		err  error
	)

	select {
	case err = <-errChan:
		return "", err
	case temp = <-fileChan:
		job.Temp = temp
		defer temp.Close()
		defer os.Remove(temp.Name())
	}

	name := temp.Name()
	//stats, _ := temp.Stat()
	//fmt.Printf("%+v\n%+v \n\n", name, stats)

	return name, err
}
