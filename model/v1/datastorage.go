package model

// Data defines a table in sql db, value in k-v pair in no sql db
type Data = map[string]interface{}

// SetInput defines
type SetInput struct {
	document string
	data     Data
	where    Data
}

// GetInput defines
type GetInput struct {
	document string
	column   []string
	where    Data
}

// IncInput defines
type IncInput struct {
	document string
	column   []string
	where    Data
}

// DBSource defines the data source
type DBSource interface {
	Set(set SetInput) error
	Get(get GetInput) (Data, error)
	Inc(inc IncInput) error
}

// DBCache defines a cache for db
type DBCache interface {
	SetCache(key string, value Data) error
	GetCache(key string) (Data, error)
	DeleteCache(key interface{}) error
}

// CachedDB defines data storage for all the service
// db write operate db directly
// db read search cache(memory / redis ...) first, if not, read db
type CachedDB struct {
	source DBSource
	cache  DBCache
}
