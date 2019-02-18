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
	Set(document string, data Data, where Data) error // update & insert
	Get(document string, where Data) ([]Data, error)
	GetOne(document string, where Data) (Data, error)
}

// DBCache defines a cache for db
type DBCache interface {
	SetCache(key string, value Data) error
	GetCache(key string) (Data, error)
	DeleteCache(key interface{}) error
}

// CachedDB defines data storage
type CachedDB interface {
	NoCache() DBDriver                                // NoCache operate db directly
	Set(document string, data Data, where Data) error // update & insert
	Get(document string, where Data) ([]Data, error)
	GetOne(document string, where Data) (Data, error)
}
