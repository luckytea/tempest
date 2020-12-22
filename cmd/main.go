package main

import (
	"fmt"

	"github.com/luckytea/tempest/cfg"
)

var version = "dev"

func main() {
	timeseries := "http_requests_total"
	labels := ""
	config := cfg.Init()

	fmt.Println("service launched:", config.Version)
}

// TODO:
// 1. hardcode and launch first
// 2. launch with args only, no config
// 3. check how timestamping works with prom library
// 3. no values, generate seed for timeseries and increment with some algorithm

// go get github.com/prometheus/client_golang/prometheus
// go get github.com/prometheus/client_golang/prometheus/promauto
// go get github.com/prometheus/client_golang/prometheus/promhttp
