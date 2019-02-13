package model

import (
	"github.com/golang/protobuf/proto"
)

// Handler describes service handler
type Handler func(req proto.Message) proto.Message

// BinaryMessage describes websocket binary message
type BinaryMessage struct {
	ProtoID uint16
	Body    []byte
}
