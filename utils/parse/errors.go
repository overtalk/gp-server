package parse

import (
	"fmt"
)

// Err : all
type Err struct {
	method string
	origin interface{}
}

func (e Err) Error() string {
	return fmt.Sprintf("failed to parse [%v] to %s", e.origin, e.method)
}
