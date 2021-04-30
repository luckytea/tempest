package generator_test

import (
	"errors"
	"testing"

	"github.com/luckytea/tempest/generator"
	"github.com/stretchr/testify/assert"
)

func TestGenerateMetricsShouldFailWithUnsupportedMetricType(t *testing.T) {
	// arrange
	var (
		expectedError error = generator.ErrUnsupportedMetricType
		ts                  = generator.Timeseries{
			Name:       "http_total_ups_and_down",
			Labels:     nil,
			MetricType: "gauge",
		}
	)

	// act
	var result, err = generator.GenerateMetric(ts)

	// assert
	t.Run("simple metrics: success", func(t *testing.T) {
		assert.Nil(t, result)
		assert.True(t,
			errors.Is(err, expectedError),
		)
	})
}
