// Package cfg contains the main configuration of the application.
package cfg

import (
	"flag"

	"github.com/LuckyTea/tempest/model"
)

func Init() *Config {
	cfg := &Config{}

	flag.StringVar(
		&cfg.Name,
		"name", "",
		"The desired name for the metric.",
	)

	flag.StringVar(
		&cfg.Desc,
		"desc", "",
		"Metric description.",
	)

	flag.StringVar(
		&cfg.Type,
		"type", "",
		"Metric type (counter).",
	)

	flag.Parse()

	return cfg
}

func (c *Config) Validate() (*Config, error) {
	switch c.Type {
	case "counter":
	default:
		return nil, model.ErrUnsupportedMetricType
	}

	return c, nil
}
