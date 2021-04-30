package generator

import "errors"

var ErrUnsupportedMetricType = errors.New("metric type unsupported")

const CounterType = "counter"

type Timeseries struct {
	Name       string
	Labels     map[string]string
	MetricType string
}
type MetricHelper struct {
	Help string
	Type string
}
