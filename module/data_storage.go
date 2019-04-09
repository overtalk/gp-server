package module

import (
	"sync"
)

// DataStorage : all data storage system for the system
type DataStorage struct {
	// JudgeFilePath string
	DB      DB
	Cache   Cache
	Configs sync.Map
}
