package utils

import "encoding/json"

func ObjectToJsonString(object interface{}) (string, error) {
	bytes, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
