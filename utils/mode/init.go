package mode

import (
	"github.com/QHasaki/gp-server/logger"
)

const (
	// DevMode means develop mode(default mode)
	DevMode int64 = 0
	// DeployMode means deploy mode
	DeployMode int64 = 1
	// TestMode means test, skip authentication
	TestMode int64 = 2
)

// Mode mark running mode of the server (default : DeployMode)
var Mode int64 = DeployMode

// SetMode set mode for the server
func SetMode(mode int64) {
	logger.Sugar.Infof("Now Runing with Mode : %s", GetModeName(mode))
	Mode = mode
}

// GetMode get mode of the server
func GetMode() int64 {
	return Mode
}

// GetModeName returns basic introduction of each mode
func GetModeName(mode int64) string {
	switch mode {
	case DevMode:
		return "Dev Mode"
	case DeployMode:
		return "Deploy Mode"
	case TestMode:
		return "Test Mode"
	default:
		return "Warning : Unknown Mode"
	}
}
