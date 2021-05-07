# tempest [![Go Reference](https://pkg.go.dev/badge/luckytea/tempest.svg)](https://pkg.go.dev/github.com/luckytea/tempest) [![Go](https://github.com/luckytea/tempest/actions/workflows/go.yml/badge.svg)](https://github.com/luckytea/tempest/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/luckytea/tempest)](https://goreportcard.com/report/github.com/luckytea/tempest)

Supreme prometheus metrics backfiller.

## Run

```sh
tempest -type=counter -name=http_requests_total -desc="Total requests" -label="code,200,3;code,404,5" -from=1620388800 -to=1620388900
```

## Supported metric types

* counter

## Parameters

| Command line | Default | Description                            |
|--------------|---------|----------------------------------------|
| name         |         | The desired name for the metric        |
| desc         |         | Metric description                     |
| type         |         | Metric type                            |
| label        |         | Label type string [name,value,counter] |
| from         |         | Time start                             |
| to           |         | Time end                               |
