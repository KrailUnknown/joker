package utils

import "encoding/json"

func GetByKeyFromJson(j []byte, key string) (string, error) {
	c := make(map[string]json.RawMessage)

	if err := json.Unmarshal(j, &c); err != nil {
		return "", err
	}

	return string(c[key]), nil
}
