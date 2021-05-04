// Package generator contains methods for generating metrics.
package generator

import (
	"fmt"

	"github.com/prometheus/prometheus/pkg/labels"
)

func EmitOpenMetrics(metric Timeseries) (string, error) {
	switch metric.MetricType {
	case CounterType:
		f := fmt.Sprintf("%#v", labels.FromStrings("__name__", "name_here", "code", "400"))
		return f, nil
	default:
		return "", ErrUnsupportedMetricType
	}
}

func OpenMetricsLine(metric Timeseries) string {
	line := fmt.Sprintf(helpTemplate, metric.Name)

	x := fmt.Sprintf(typeTemplate, metric.Name, metric.MetricType)

	line += x

	for i := range metric.Samples {
		x := fmt.Sprintf(metricTemplate, metric.Name, "label_key", "label_value",
			metric.Samples[i].Value, metric.Samples[i].Timestamp)
		line += x
	}

	line = fmt.Sprintf("%s%s", line,
		eofTemplate)

	return line
}
