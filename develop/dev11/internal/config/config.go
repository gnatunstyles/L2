package config

type Config struct {
	Address string
}

func New() *Config {
	var c Config
	c.Address = "localhost:8080"
	return &c
}
