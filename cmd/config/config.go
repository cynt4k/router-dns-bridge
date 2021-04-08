package config

var (
	config = &Config{}
)

type Config struct {
	DevMode bool `yaml:"devMode"`
}

func GetConfig() *Config {
	return config
}
