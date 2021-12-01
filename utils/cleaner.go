package utils

import "encoding/json"

func PrettyStruct(object interface{}) (string, error) {
	b, err := json.MarshalIndent(object, "", "  ")

	if err != nil {
		return "", err
	}

	return string(b), nil
}

//ScreenName
