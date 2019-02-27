package test

import (
	"github.com/qinhan-shu/gp-server/protocol"
)

// Test is the
func Test(req *protocol.TestRequest) *protocol.TestResponse {
	return &protocol.TestResponse{
		A: "111",
		B: "111",
	}
}
