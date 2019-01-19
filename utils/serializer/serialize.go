package serializer

// Encode data with the specified serializer
func Encode(v interface{}) ([]byte, error) {
	return jsonMarshal(v)
}

// Decode data with the specified serializer
func Decode(data []byte, v interface{}) error {
	protoID, err := getSerializerProtoID(data)
	if err != nil {
		return err
	}
	switch protoID {
	case json_serializer:
		return jsonUnmarshal(data[protoLen:], v)
	default:
		return ErrInValidSerializer
	}
}
