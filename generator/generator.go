package generator

import "fmt"

func OpenMetricsLine(metric Timeseries) string {
	var header = generateHeader(metric)

	line := header

	for i := range metric.Samples {
		x := fmt.Sprintf(metricTemplate,
			metric.Name, metric.Samples[i].LabelName, metric.Samples[i].LabelValue,
			metric.Samples[i].Value, metric.Samples[i].Timestamp)
		line += x
	}

	line = fmt.Sprintf("%s%s", line, eofTemplate)

	return line
}

func generateHeader(metric Timeseries) string {
	var header = fmt.Sprintf(helpTemplate, metric.Name, metric.Desc)

	header += fmt.Sprintf(typeTemplate, metric.Name, metric.MetricType)

	return header
}
