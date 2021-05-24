package julenkv

const (
	DefaultAddr = "localhost:5200"
)

type Config struct {
	Addr string `json:"addr"`
}

func DefaultConfig() Config {
	return Config{
		Addr: DefaultAddr,
	}
}

