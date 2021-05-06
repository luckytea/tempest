package generator_test

import (
	"testing"

	"github.com/luckytea/tempest/generator"

	"github.com/stretchr/testify/require"
)

var result string

func Benchmark_OpenMetricsLine(b *testing.B) {
	var metric = generator.Timeseries{
		MetricType: generator.CounterType,
		Name:       "http_total_requests",
		Desc:       "The total number of HTTP requests.",
		Samples:    []generator.BackfillSample{},
	}

	var r string

	for n := 0; n < b.N; n++ {
		r = generator.OpenMetricsLine(metric)
	}

	result = r
}

func Test_OpenMetricsLine_Success(t *testing.T) {
	t.Run("OpenMetricsLine: success", func(t *testing.T) {
		// arrange
		var (
			metric = generator.Timeseries{
				MetricType: generator.CounterType,
				Name:       "http_total_requests",
				Desc:       "The total number of HTTP requests.",
				Samples:    []generator.BackfillSample{},
			}

			want string = `# HELP http_total_requests The total number of HTTP requests.
# TYPE http_total_requests counter
# EOF
`
		)

		// act
		got := generator.OpenMetricsLine(metric)

		// assert
		require.Equal(t,
			want, got,
		)
	})
}
