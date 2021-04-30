package generator

import (
	"github.com/prometheus/client_golang/prometheus"
)

func GenerateMetric(metric Timeseries) (interface{}, error) {
	switch metric.MetricType {
	case CounterType:
		return prometheus.CounterVec{}, nil
	default:
		return nil, ErrUnsupportedMetricType
	}
}

func EmitOpenMetrics(metric interface{}) (string, error) {
	switch metric.(type) {
	case prometheus.CounterVec:
		return "success", nil
	default:
		return "", ErrUnsupportedMetricType
	}
}
