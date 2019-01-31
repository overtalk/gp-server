package gate

import (
	"sync"

	"github.com/QHasaki/Server/model/v1"
)

// ServiceHandler maps proto id to handler
type ServiceHandler struct {
	protoMap map[uint16]model.Handler
	l        *sync.RWMutex
}
