// Package generator contains methods for generating metrics.
package generator

import "fmt"

func OpenMetricsLine(metric Timeseries) string {
	var header = generateHeader(metric)

	line := header

	for i := range metric.Samples {
		x := fmt.Sprintf(metricTemplate, metric.Name, "label_key", "label_value",
			metric.Samples[i].Value, metric.Samples[i].Timestamp)
		line += x
	}

	line = fmt.Sprintf("%s%s", line,
		eofTemplate)

	return line
}

func generateHeader(metric Timeseries) string {
	var header = fmt.Sprintf(helpTemplate, metric.Name, metric.Desc)

	header += fmt.Sprintf(typeTemplate, metric.Name, metric.MetricType)

	return header
}
