package test

import (
	"github.com/QHasaki/Server/protocol/v1"
)

// Test is the
func Test(req *protocol.TestRequest) *protocol.TestResponse {
	return &protocol.TestResponse{
		A: "111",
		B: "111",
	}
}
