package model

// DataStorage defines data storage model
type DataStorage struct {
}

// Config defines server config
type Config interface {
	GetDataStorage() (*DataStorage, error)
}
