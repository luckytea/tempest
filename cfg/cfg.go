// Package cfg contains the main configuration of the application.
package cfg

import (
	"flag"

	"github.com/luckytea/tempest/model"
)

func Init() *Config {
	cfg := &Config{}

	flag.StringVar(
		&cfg.Type,
		"type", "",
		"Metric type [counter].",
	)

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
		&cfg.Label,
		"label", "",
		"Metric label, format: name,value,count",
	)

	flag.Int64Var(
		&cfg.From,
		"from", 0,
		"Metric sample from unixtime.",
	)

	flag.Int64Var(
		&cfg.To,
		"to", 0,
		"Metric sample to unixtime.",
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

	if c.From >= c.To {
		return nil, ErrMalformedTime
	}

	return c, nil
}
