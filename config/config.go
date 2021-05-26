package config

const (
	DefaultAddr = "0.0.0.0:5200"
)

type Config struct {
	Addr string `json:"addr"`
}

func DefaultConfig() Config {
	return Config{
		Addr: DefaultAddr,
	}
}

