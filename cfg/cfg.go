package cfg

import (
	"flag"
	"strings"
)

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

	flag.StringVar(
		&cfg.Ts.LabelsLine,
		"labels", "label,label1",
		"Metric labels.",
	)

	for _, label := range strings.Split(cfg.Ts.LabelsLine, ",") {
		cfg.Ts.Labels = append(cfg.Ts.Labels, Label{Name: label, Values: []string{"pupa", "lupa", "zalupa"}})
	}

	flag.StringVar(
		&cfg.Ts.MetricType,
		"type", "counter",
		"Metric type(counter, gauge)",
	)

	flag.StringVar(
		&cfg.Ts.Name,
		"name", "default_metric_name_counter_total",
		"Metric name.",
	)

	return cfg
}
