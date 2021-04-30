package cfg

import (
	"flag"
)

func Init(v string) *Config {
	cfg := &Config{}

	flag.StringVar(
		&cfg.Version,
		"version", v,
		"Service version.",
	)

	flag.StringVar(
		&cfg.Ts.LabelsLine,
		"labels", "label,label2,label3",
		"Metric labels.",
	)

	// for _, label := range strings.Split(cfg.Ts.LabelsLine, ",") {
	// 	cfg.Ts.Labels = append(cfg.Ts.Labels, Label{Name: label, Values: []string{"pupa", "lupa", "zalupa"}})
	// }

	cfg.Ts.Labels = append(cfg.Ts.Labels, Label{Name: "status", Values: []string{"200", "404", "502"}})

	// cfg.Ts.Labels = append(cfg.Ts.Labels, Label{Name: "method", Values: []string{"GET", "POST", "DELETE"}})
	// cfg.Ts.Labels = append(cfg.Ts.Labels, Label{Name: "handler", Values: []string{"order", "product"}})

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
