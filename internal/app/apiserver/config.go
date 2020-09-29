package apiserver

type Config struct {
	BindAddr string `json:"bind_addr"`
	LogLev   string `json:"log_lev"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLev:   "debug",
	}
}
