package main

import (
	"fmt"
	"os"

	"github.com/luckytea/tempest/cfg"

	"github.com/prometheus/client_golang/prometheus"
)

// TODO:
// 1. hardcode and launch first +
// 2. launch with args only, no config +
// 3. check how timestamping works with prom library
// set label values
// mapping from provided metric type to prom lib metric type
// remove hardcoded values
// 3. no values, generate seed for timeseries and increment with some algorithm

var version = "dev"

func main() {
	config := cfg.Init(version)

	fmt.Println("service launched:", config.Version, "at port", config.Port)

	switch config.Ts.MetricType {
	case "counter":
	default:
		fmt.Println("unsupported metric type")
		os.Exit(1)
	}

	metric := generateCounterMetric(&config.Ts)

	fmt.Println(metric)

	arr := generateMatrix(config.Ts.Labels)

	for i := range arr {
		fmt.Println(arr[i])
	}

	os.Exit(0)

}

func extractNames(l []cfg.Label) []string {
	var labels = make([]string, len(l))

	for i := range l {
		labels[i] = l[i].Name
	}

	return labels
}

func generateCounterMetric(t *cfg.Timeseries) *prometheus.CounterVec {
	// generate openmetric-formatted metric from cli

	labels := extractNames(t.Labels)

	var metric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: t.Name,
		}, labels)

	prometheus.MustRegister(metric)

	return metric
}

func generateMatrix(labels []cfg.Label) [][]string {
	var (
		tags = make([]string, len(labels))
		some [][]string
	)

	for i := range tags {
		tags[i] = labels[i].Values[0]
	}

	some = append(some, tags)

	return some
}
