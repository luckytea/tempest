package generator

import (
	"errors"
)

var ErrUnsupportedMetricType = errors.New("metric type unsupported")

const CounterType = "counter"

type Timeseries struct {
	MetricType string
	Name       string
	Samples    []BackfillSample
}

type BackfillSample struct {
	Timestamp int64
	Value     float64
	Labels    map[string]string
}
