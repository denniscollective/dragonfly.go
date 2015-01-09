package dragonfly

import (
	//"fmt"
	"io/ioutil"
	"os"
)

type Step struct {
	Args    []string
	Command string
}

func (step Step) Fetch(fileChan chan *os.File, errChan chan error) {
	filename := step.Args[0]
	temp, err := fechFile(filename)

	if err != nil {
		errChan <- err
		return
	}

	fileChan <- temp

}

func fechFile(filename string) (*os.File, error) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	temp, err := ioutil.TempFile(os.TempDir(), "godragonfly")

	if err != nil {
		panic(err)
	}

	_, err = temp.Write(content)
	if err != nil {
		panic(err)
	}

	return temp, err
}
