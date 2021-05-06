package generator

import "errors"

var ErrUnsupportedMetricType = errors.New("metric type unsupported")

const CounterType = "counter"

type Timeseries struct {
	MetricType string
	Name       string
	Desc       string
	Samples    []BackfillSample
}

type BackfillSample struct {
	Timestamp int64
	Value     float64
	Labels    map[string]string
}

const (
	helpTemplate   string = "# HELP %s %s\n"
	typeTemplate   string = "# TYPE %s %s\n"
	eofTemplate    string = "# EOF\n"
	metricTemplate string = `%s{%s="%s"} %v %v`
)
