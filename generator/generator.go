package generator

import "github.com/prometheus/client_golang/prometheus"

func EmitOpenMetrics(metric interface{}) (string, error) {
	switch metric.(type) {
	case prometheus.CounterVec:
		return "success", nil
	default:
		return "", ErrUnsupportedMetricType
	}
}
