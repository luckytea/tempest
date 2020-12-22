package cfg

type Config struct {
	Version string
	Port    string
	Host    string
	Ts      Timeseries
}

type Timeseries struct {
	Name       string
	LabelsLine string
	Labels     []Label
	MetricType string
}

type Label struct {
	Name   string
	Values []string
}
