package dragonfly

import (
	//"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type Step interface {
	Apply(in chan *os.File, errIn chan error) (out chan *os.File, errChan chan error)
	//Args    []string
	//Command string
}

type Job struct {
	Steps []Step
	Temp  *os.File
}

func (job *Job) Apply() (temp *os.File, err error) {
	tail := make(chan *os.File)
	errChan := make(chan error)
	close(tail)
	defer close(errChan)

	for _, step := range job.Steps {
		tail, errChan = step.Apply(tail, errChan)
	}

	select {
	case err = <-errChan:
	case temp = <-tail:
	}

	return
}

type stepApplication func(temp *os.File) (*os.File, error)

func applyStepPipeline(in chan *os.File, errIn chan error, step stepApplication) (out chan *os.File, errChan chan error) {
	out = make(chan *os.File)
	errChan = make(chan error)

	go func() {
		defer close(out)
		defer close(errChan)

		var (
			err     error
			content *os.File
		)

		select {
		case prev := <-in:
			content, err = step(prev)

		case err = <-errIn:
		}

		if err != nil {
			errChan <- err
			return
		}

		out <- content

	}()

	return out, errChan
}

type FetchFileStep struct {
	Args    []string
	Command string
}

type ResizeStep struct {
	Args    []string
	Command string
}

func (step ResizeStep) Apply(in chan *os.File, errIn chan error) (out chan *os.File, errChan chan error) {
	return applyStepPipeline(in, errIn, func(temp *os.File) (newTemp *os.File, err error) {
		format := step.Args[1]
		return step.resize(temp, format)
	})
}

func (step FetchFileStep) Apply(in chan *os.File, errIn chan error) (out chan *os.File, errChan chan error) {
	return applyStepPipeline(in, errIn, func(_ *os.File) (temp *os.File, err error) {
		filename := step.Args[0]
		return fechFile(filename)
		//return nil, errors.New("please don't stop the music")
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

	if image == nil {
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
