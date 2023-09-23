package handlers

import "encoding/json"

func ErrToJson(err error) []byte {
	errJSON := map[string]string{
		"message": err.Error(),
	}

	jsonData, err := json.Marshal(errJSON)
	if err != nil {
		return []byte(`{"message": "there was an issue marshaling data to json"}`)
	}

	return jsonData
}

func StringToJson(message string) []byte {
	errJSON := map[string]string{
		"message": message,
	}

	jsonData, err := json.Marshal(errJSON)
	if err != nil {
		return []byte(`{"message": "there was an issue marshaling data to json"}`)
	}

	return jsonData
}
