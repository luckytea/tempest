package model

type MetricHelper struct {
	Help string
	Type string
}

type Config struct {
	Version string
	Ts      Timeseries
}

type Timeseries struct {
	Name       string
	Labels     []Label
	MetricType string
}

type Label struct {
	Name   string
	Values []string
}
