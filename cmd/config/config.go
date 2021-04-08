package config

var (
	config = &Config{}
)

type RouterOptions struct {
	AllowedIPs []string `yaml:"allowedIps"`
}
type Router struct {
	APIKey   string        `yaml:"apiKey"`
	Domain   string        `yaml:"domain"`
	Provider string        `yaml:"provider"`
	Options  RouterOptions `yaml:"routerOptions"`
}
type ProviderPowerdns struct {
	URL    string `yaml:"url"`
	APIKey string `yaml:"apiKey"`
}
type Providers struct {
	Powerdns ProviderPowerdns `yaml:"powerdns"`
}
type API struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Config struct {
	DevMode   bool              `yaml:"devMode"`
	API       API               `yaml:"api"`
	Providers Providers         `yaml:"providers"`
	Routers   map[string]Router `yaml:"routers"`
}

func GetConfig() *Config {
	return config
}
