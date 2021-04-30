package generator

import "errors"

var ErrUnsupportedMetricType = errors.New("metric type unsupported")

const CounterType = "counter"
