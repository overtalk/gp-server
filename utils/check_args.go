package utils

import (
	"fmt"
)

// CheckArgs : check the keys in fields is in the map
func CheckArgs(args map[string]interface{}, fields ...string) error {
	for _, key := range fields {
		if _, ok := args[key]; !ok {
			return fmt.Errorf("key [%s] is missing", key)
		}
	}
	return nil
}
