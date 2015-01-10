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
	fmt.Println(jobStr)

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

func decodeJobStr(b64Str string) ([]byte, error) {
	b64Str += "=" //dragonfly trims a trailing =\n from the jobs
	jobStr, err := base64.StdEncoding.DecodeString(b64Str)

	if err != nil {
		fmt.Println(b64Str)
		b64Str += "="
		fmt.Println(b64Str)

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
