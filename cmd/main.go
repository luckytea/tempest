package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/luckytea/tempest/cfg"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TODO:
// 1. hardcode and launch first +
// 2. launch with args only, no config +
// 3. check how timestamping works with prom library
// set label values
// mapping from provided metric type to prom lib metric type
// remove hardcoded values
// 3. no values, generate seed for timeseries and increment with some algorithm

// go get github.com/prometheus/client_golang/prometheus
// go get github.com/prometheus/client_golang/prometheus/promauto
// go get github.com/prometheus/client_golang/prometheus/promhttp

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

	// http_requests_total{code,handler,service,method}
	// for every method for every handler for every code for every service
	// WithLabelValues{(1,2,3,4)}.inc()

	//

	// for every key:
	// dict[key] = value
	// unordereddict[key] = value
	metric := generateCounterMetric(&config.Ts)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe("localhost"+config.Port, nil))
	}()

	go func() {
		for {
			var wg sync.WaitGroup

			for i := range config.Ts.Labels {
				wg.Add(1)

				go mutate(&wg, metric, config.Ts.Labels[i], i)
			}

			wg.Wait()

			time.Sleep(1 * time.Second)
		}
	}()

	// waiting for listener
	time.Sleep(50 * time.Millisecond)

	ossig := make(chan os.Signal, 1)

	signal.Notify(ossig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-ossig
}

func extractNames(l []cfg.Label) []string {
	var labels = make([]string, len(l))

	for i := range l {
		labels[i] = l[i].Name
	}

	return labels
}

func generateCounterMetric(t *cfg.Timeseries) *prometheus.CounterVec {
	labels := extractNames(t.Labels)

	var metric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: t.Name,
		}, labels)

	prometheus.MustRegister(metric)

	return metric
}

func mutate(wg *sync.WaitGroup, metric *prometheus.CounterVec, label cfg.Label, i int) {
	metric.WithLabelValues(label.Values[i], label.Values[i]).Inc()
	wg.Done()
}

func generateMatrix([]cfg.Label) [][]string {


	return nil
}
