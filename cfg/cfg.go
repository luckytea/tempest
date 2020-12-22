package cfg

import "flag"

func Init(v string) *Config {
	cfg := &Config{}

	flag.StringVar(
		&cfg.Version,
		"version", v,
		"Service version.",
	)

	flag.StringVar(
		&cfg.Port,
		"port", ":10500",
		"Service port.",
	)

	return cfg
}
