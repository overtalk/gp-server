package model

// Where defines the query condition
// Operator(>, <, =, <=, >=, !=)
// Value means the details
type Where struct {
	Operator string
	Value    interface{}
}

// Data defines a table in sql db, value in k-v pair in no sql db
type Data = map[string]interface{}

// DBDriver defines the data source
type DBDriver interface {
	Set(document string, data Data, where Data) error
	Get(document string, column []string, where Data) (Data, error)
	Inc(document string, column []string, where Data) error
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
	source DBDriver
	cache  DBCache
}
