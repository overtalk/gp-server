package utils

import (
	"errors"
	"math/rand"
	"time"
)

var (
	// ErrIllegalRange means minimum > maximum
	ErrIllegalRange = errors.New("illegal random number range [ minimum > maximum ]")
)

// RandInt is to get a rand int
// return min ~ max
func RandInt(min, max int) (int, error) {
	if min > max {
		return 0, ErrIllegalRange
	}
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min+1) + min
	return randNum, nil
}
