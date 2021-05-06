package generator_test

import (
	"testing"

	"github.com/luckytea/tempest/generator"

	"github.com/stretchr/testify/require"
)

func Benchmark_OpenMetricsLine(b *testing.B) {
	var metric = generator.Timeseries{
		MetricType: generator.CounterType,
		Name:       "http_total_requests",
		Desc:       "The total number of HTTP requests.",
		Samples:    []generator.BackfillSample{},
	}

	for n := 0; n < b.N; n++ {
		_ = generator.OpenMetricsLine(metric)
	}
}

func Test_OpenMetricsLine_Success(t *testing.T) {
	t.Run("OpenMetricsLine: success", func(t *testing.T) {
		// arrange
		var (
			metric = generator.Timeseries{
				MetricType: generator.CounterType,
				Name:       "http_requests_total",
				Desc:       "The total number of HTTP requests.",
				Samples: []generator.BackfillSample{
					{
						Timestamp:  1620388800,
						Value:      1,
						LabelName:  "code",
						LabelValue: "200",
					},
				},
			}

			want string = `# HELP http_requests_total The total number of HTTP requests.
# TYPE http_requests_total counter
http_requests_total{code="200"} 1 1620388800
# EOF`
		)

		// act
		got := generator.OpenMetricsLine(metric)

		// assert
		require.Equal(t,
			want, got,
		)
	})
}
