package generator

import (
	"github.com/luckytea/tempest/model"
	"github.com/prometheus/client_golang/prometheus"
)

func GenerateMetric(metric model.Timeseries) (interface{}, error) {
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
