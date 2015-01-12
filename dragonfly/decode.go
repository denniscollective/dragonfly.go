package dragonfly

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func Decode(str string) (*Job, error) {
	var job Job

	jobStr, err := decodeJobStr(str)
	if err != nil {
		return &job, err
	}

	jobArr, err := decodeJson(jobStr)
	if err != nil {
		return &job, err
	}

	job.Steps = make([]Step, len(jobArr))

	for i, v := range jobArr {
		job.Steps[i] = StepFromArray(v)
	}

	return &job, err
}

func StepFromArray(array []string) Step {
	switch array[0] {
	case "ff":
		return &FetchFileStep{Command: array[0], Args: array[1:]}
	case "p":
		return &ResizeStep{Command: array[0], Args: array[1:]}

	}

	return nil
}

func decodeJobStr(b64Str string) ([]byte, error) {
	b64Str += "=" //dragonfly trims a trailing =\n from the jobs
	jobStr, err := base64.StdEncoding.DecodeString(b64Str)

	if err != nil {
		b64Str += "="
		jobStr, err2 := base64.StdEncoding.DecodeString(b64Str)
		err = err2

		if err2 != nil {
			fmt.Println("\n******************************************************************************")
			fmt.Println(b64Str)
			fmt.Println(jobStr)
			fmt.Println("******************************************************************************\n")
		}
	}

	return jobStr, err
}

func decodeJson(jsonStr []byte) ([][]string, error) {
	var i [][]string
	err := json.Unmarshal(jsonStr, &i)
	if err != nil {
		str := string(jsonStr) + "]"
		err = json.Unmarshal([]byte(str), &i)

	}
	return i, err
}
