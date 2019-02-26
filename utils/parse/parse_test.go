package parse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/gp-server/utils/parse"
)

func TestBytes(t *testing.T) {
	var in interface{}
	var out []byte

	// []byte
	in = []byte("qinhan")
	out = parse.Bytes(in)
	if !assert.Equal(t, in, out) {
		t.Errorf("Unexpected result")
		return
	}

	// string
	in = "qinhan"
	out = parse.Bytes(in)
	if !assert.Equal(t, []byte("qinhan"), out) {
		t.Errorf("Unexpected result")
		return
	}

	// nil
	in = nil
	out = parse.Bytes(in)
	if !assert.Equal(t, []byte(nil), out) {
		t.Errorf("Unexpected result")
		return
	}

	// others
	in = []int64{0, 1, 2}
	out = parse.Bytes(in)
	if !assert.Equal(t, []byte(nil), out) {
		t.Errorf("Unexpected result")
		return
	}

}

func TestInt(t *testing.T) {
	var in interface{}
	var out int64

	// string
	in = "12"
	out = parse.Int(in)
	if !assert.Equal(t, int64(12), out) {
		t.Errorf("Unexpected result")
		return
	}

	// error string
	in = "1a2"
	out = parse.Int(in)
	if !assert.Equal(t, int64(0), out) {
		t.Errorf("Unexpected result")
		return
	}

	// uint16
	in = uint16(90)
	out = parse.Int(in)
	if !assert.Equal(t, int64(90), out) {
		t.Errorf("Unexpected result")
		return
	}
	// uint32
	in = uint32(90)
	out = parse.Int(in)
	if !assert.Equal(t, int64(90), out) {
		t.Errorf("Unexpected result")
		return
	}
}

func TestString(t *testing.T) {
	var in interface{}
	var out string

	// string
	in = "qinhan"
	out = parse.String(in)
	if !assert.Equal(t, "qinhan", out) {
		t.Errorf("Unexpected result")
		return
	}

	// []uint8 ( []byte )
	in = []uint8("qinhan")
	out = parse.String(in)
	if !assert.Equal(t, "qinhan", out) {
		t.Errorf("Unexpected result")
		return
	}
}
