# tempest [![Go Reference](https://pkg.go.dev/badge/luckytea/tempest.svg)](https://pkg.go.dev/github.com/luckytea/tempest) [![Go](https://github.com/luckytea/tempest/actions/workflows/go.yml/badge.svg)](https://github.com/luckytea/tempest/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/luckytea/tempest)](https://goreportcard.com/report/github.com/luckytea/tempest)

Supreme prometheus metrics backfiller.

## Run

```sh
tempest -type=counter -name=metric_name -desc=metric_desc
```

## Supported metric types

* counter

## Parameters

| Command line | Default | Description |
|--------------|---------|-------------|
| name         |         | The desired name for the metric. |
| desc         |         | Metric description. |
| type         |         | Metric type.        |
