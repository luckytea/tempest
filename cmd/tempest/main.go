package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LuckyTea/tempest/cfg"
	"github.com/LuckyTea/tempest/generator"
)

func main() {
	config, err := cfg.Init().Validate()
	if err != nil {
		log.Println("can't initiate application cause:", err)
		os.Exit(1)
	}

	var m = generator.Timeseries{
		MetricType: config.Type,
		Name:       config.Name,
		Desc:       config.Desc,
	}

	fmt.Println(generator.OpenMetricsLine(m))
}
