package main

import (
	"fmt"
	"log"
	"os"

	"github.com/luckytea/tempest/cfg"
	"github.com/luckytea/tempest/generator"
)

func main() {
	config, err := cfg.Init().Validate()
	if err != nil {
		log.Println("can't initiate application cause:", err)
		os.Exit(1)
	}

	samples, err := generator.GenerateSamplesFromString(config.Label, config.From, config.To)
	if err != nil {
		log.Println("can't generate labels:", err)
		os.Exit(1)
	}

	var m = generator.Timeseries{
		MetricType: config.Type,
		Name:       config.Name,
		Desc:       config.Desc,
		Samples:    samples,
	}

	fmt.Println(generator.OpenMetricsLine(m))
}
