package serializer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/utils/serializer"
)

type message struct {
	ID   int64
	Data string
}

func Test_Serialize(t *testing.T) {
	m1 := &message{
		ID:   100,
		Data: "Hello world!",
	}
	data, err := serializer.Encode(m1)
	if err != nil {
		t.Errorf("failed to encode: %v", err)
		return
	}
	m2 := new(message)
	if err = serializer.Decode(data, m2); err != nil {
		t.Errorf("failed to decode: %v", err)
		return
	}
	if !assert.Equal(t, m1, m2) {
		t.Errorf("m1[%v] != m2[%v]", m1, m2)
		return
	}
}
