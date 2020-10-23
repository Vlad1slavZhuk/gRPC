package json

import "encoding/json"

func Encode(i interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func Decode(b []byte, i interface{}) (interface{}, error) {
	//TODO

	return nil, nil
}
