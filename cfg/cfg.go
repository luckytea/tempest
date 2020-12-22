package cfg

import "flag"

func Init(v string) *Config {
	cfg := &Config{}

	flag.StringVar(
		&cfg.Version,
		"version", v,
		"Service version",
	)

	return cfg
}
