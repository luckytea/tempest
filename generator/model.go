package generator

import "errors"

var (
	ErrUnsupportedMetricType = errors.New("metric type unsupported")
	ErrMalformed             = errors.New("malformed label string")
)

const CounterType = "counter"

type Timeseries struct {
	MetricType string
	Name       string
	Desc       string
	Samples    []BackfillSample
}

type BackfillSample struct {
	Timestamp  int64
	Value      float64
	LabelName  string
	LabelValue string
}

const (
	helpTemplate   string = "# HELP %s %s\n"
	typeTemplate   string = "# TYPE %s %s\n"
	metricTemplate string = "%s{%s=\"%s\"} %v %d\n"
	eofTemplate    string = "# EOF"
)
