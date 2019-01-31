package config

type Config map[string]string

// NewConfig is the constructor of Config
func NewConfig() Config {
	config := make(Config)
	return config
}
