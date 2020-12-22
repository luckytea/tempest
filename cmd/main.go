package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/luckytea/tempest/cfg"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var version = "dev"

type Timeseries struct {
	name        string
	labels      string
	metric_type string
}

func main() {
	config := cfg.Init(version)

	fmt.Println("service launched:", config.Version, "at port", config.Port)

	t := &Timeseries{
		name:        "http_requests_total",
		labels:      "service,code,handler",
		metric_type: "counter",
	}

	metric := generateCounterMetric(t)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe("localhost"+config.Port, nil))
	}()

	fmt.Println("inc")

	metric.WithLabelValues("service", "code", "handler").Inc()

	time.Sleep(50 * time.Millisecond)

	fmt.Println("make get")

	resp, err := http.Get("http://localhost" + config.Port + "/metrics")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("read body")

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println("body:", bodyString)
}

// TODO:
// 1. hardcode and launch first
// 2. launch with args only, no config
// 3. check how timestamping works with prom library
// mapping from provided metric type to prom lib metric type
// 3. no values, generate seed for timeseries and increment with some algorithm

// go get github.com/prometheus/client_golang/prometheus
// go get github.com/prometheus/client_golang/prometheus/promauto
// go get github.com/prometheus/client_golang/prometheus/promhttp

func generateCounterMetric(t *Timeseries) *prometheus.CounterVec {
	var metric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: t.name,
		}, strings.Split(t.labels, ","))

	prometheus.MustRegister(metric)

	return metric
}
