package dragonfly

import (
	"io/ioutil"
	"os"
	"os/exec"
)

type Step interface {
	Apply(in chan *os.File) (out chan *os.File, errChan chan error)
	//Args    []string
	//Command string
}

type FetchFileStep struct {
	Args    []string
	Command string
}

type ResizeStep struct {
	Args    []string
	Command string
}

type stepApplication func() (*os.File, error)

func applyStepPipeline(step stepApplication) (out chan *os.File, errChan chan error) {
	out = make(chan *os.File)
	errChan = make(chan error)

	go func() {
		defer close(out)
		defer close(errChan)

		content, err := step()
		if err != nil {
			errChan <- err
			return
		}

		out <- content
	}()

	return out, errChan
}

func (step ResizeStep) Apply(in chan *os.File) (out chan *os.File, errChan chan error) {
	return applyStepPipeline(func() (newTemp *os.File, err error) {
		temp := <-in
		format := step.Args[1]
		return step.resize(temp, format)
	})
}

func (step ResizeStep) resize(image *os.File, format string) (*os.File, error) {
	binary, err := exec.LookPath("convert")
	if err != nil {
		return nil, err
	}

	tempPrefix := "godragonfly" + format
	resized, err := ioutil.TempFile(os.TempDir(), tempPrefix)
	if err != nil {
		return nil, err
	}

	//if image == nil {
	//	return nil, err
	//}

	args := []string{
		image.Name(),
		"-resize", format,
		resized.Name(),
	}

	cmd := exec.Command(binary, args...)
	cmd.Run()

	return resized, err
}

func (step FetchFileStep) Apply(in chan *os.File) (out chan *os.File, errChan chan error) {
	return applyStepPipeline(func() (temp *os.File, err error) {
		filename := step.Args[0]
		return fechFile(filename)
		//return nil, errors.New("please don't stop the music")
	})
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
