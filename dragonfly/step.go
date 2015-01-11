package dragonfly

import (
	"io/ioutil"
	"os"
	"os/exec"
)

type Step struct {
	Args    []string
	Command string
}

func (step Step) Process(temp *os.File, fileChan chan *os.File, errChan chan error) {
	format := step.Args[1]
	newTemp, err := step.resize(temp, format)

	if err != nil {
		errChan <- err
		return
	}

	fileChan <- newTemp
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

func (step Step) resize(image *os.File, format string) (*os.File, error) {
	binary, err := exec.LookPath("convert")
	if err != nil {
		return nil, err
	}

	tempPrefix := "godragonfly" + format
	resized, err := ioutil.TempFile(os.TempDir(), tempPrefix)
	if err != nil {
		return nil, err
	}

	args := []string{
		image.Name(),
		"-resize", format,
		resized.Name(),
	}

	cmd := exec.Command(binary, args...)
	cmd.Run()

	return resized, err
}

func fechFile(filename string) (*os.File, error) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	temp, err := ioutil.TempFile(os.TempDir(), "godragonfly")

	if err != nil {
		return nil, err
	}

	_, err = temp.Write(content)
	if err != nil {
		return nil, err
	}

	return temp, err
}
