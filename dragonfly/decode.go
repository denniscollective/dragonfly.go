package dragonfly

import (
	"encoding/base64"
	"encoding/json"
)

type Step struct {
	command string
	args    []string
}

type Job []Step

func Decode(str string) (Job, error) {
	jobStr, err := decodeJobStr(str)
	if err != nil {
		return nil, err
	}

	jobArr, err := decodeJson(jobStr)
	if err != nil {
		return nil, err
	}

	job := make(Job, len(jobArr))

	for i, v := range jobArr {
		var step Step
		step.command = v[0]
		step.args = v[1:]
		job[i] = step
	}

	return job, err

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
