package serializer

import (
	"encoding/json"
)

func jsonMarshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return addSerializationProtoID(data, json_serializer), nil
}

func jsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
