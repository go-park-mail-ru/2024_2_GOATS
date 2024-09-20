package config

type Config struct {
	Listener Listener `yaml:"listener"`
}

type Listener struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type ConfigContextKey struct {
	Address string
	Port    int
}
