package model

// Handler describes service handler
type Handler func(in []byte, outChan chan<- []byte)

// BinaryMessage describes websocket binary message
type BinaryMessage struct {
	ProtoID uint16
	Body    []byte
}
