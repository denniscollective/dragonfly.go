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

func Decode(str string) ([]interface{}, error) {
	jobStr, err := decodeJobStr(str)
	if err != nil {
		return nil, err
	}

	jobArr, err := decodeJson(jobStr)
	if err != nil {
		return nil, err
	}

	return jobArr, err

}

func decodeJobStr(jobStr string) (*[]byte, error) {
	fixedStr := jobStr + "=\n" //dragonfly trims a trailing =\n from the jobs
	job, err := base64.StdEncoding.DecodeString(fixedStr)
	return &job, err
}

func decodeJson(str *[]byte) ([]interface{}, error) {
	var i []interface{}
	err := json.Unmarshal(*str, &i)
	return i, err
}
