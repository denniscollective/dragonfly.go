package dragonfly

import (
	"encoding/base64"
	"encoding/json"
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
		var step Step
		step.Command = v[0]
		step.Args = v[1:]
		job.Steps[i] = step
	}

	return &job, err
}

func decodeJobStr(jobStr string) (*[]byte, error) {
	fixedStr := jobStr + "=\n" //dragonfly trims a trailing =\n from the jobs
	job, err := base64.StdEncoding.DecodeString(fixedStr)
	return &job, err
}

func decodeJson(str *[]byte) ([][]string, error) {
	var i [][]string
	err := json.Unmarshal(*str, &i)
	return i, err
}
