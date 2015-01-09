package dragonfly

import (
	//"fmt"
	"os"
)

type Job struct {
	Steps []Step
	Temp  *os.File
}

func (job *Job) Apply() (string, error) {
	fileChan := make(chan *os.File)
	errChan := make(chan error)
	var (
		temp *os.File
		err  error
	)

	go job.Steps[0].Fetch(fileChan, errChan)

	select {
	case err = <-errChan:
		return "", err
	case temp = <-fileChan:
		//defer temp.Close()
		//defer os.Remove(temp.Name())
	}

	go job.Steps[1].Process(temp, fileChan, errChan)
	select {
	case err = <-errChan:
		return "", err
	case temp = <-fileChan:
		//defer temp.Close()
		//defer os.Remove(temp.Name())
	}

	name := temp.Name()
	//stats, _ := temp.Stat()
	//fmt.Printf("%+v\n%+v \n\n", name, stats)

	return name, err
}
