package generator_test

import (
	"testing"

	"github.com/luckytea/tempest/generator"
	"github.com/stretchr/testify/assert"
)

// TestEmpitUpenMetricsShouldReturnSuccess
// генерируем простую метрику в формате промифиуса и валидируем в формате open metrics.
func TestEmpitUpenMetricsShouldReturnSuccess(t *testing.T) {
	// arrange
	var (
		expectedResult string = "success"

		metric = generator.Timeseries{
			MetricType: "counter",
			Name:       "some_metrics",
			Samples: []generator.BackfillSample{
				{Value: 1, Timestamp: 2, Labels: nil},
				{Value: 3, Timestamp: 4, Labels: nil},
			},
		}

		// input prometheus.CounterVec
	)

	// act
	var result = generator.OpenMetricsLine(metric)

	// assert
	t.Run("simple metrics: success", func(t *testing.T) {
		assert.Equal(t,
			expectedResult,
			result,
		)
	})
}

// TestEmitOpenMetricsShouldFailWithUnsupportedMetricType
// передан неподдерживаемый формат метрики.
// func TestEmitOpenMetricsShouldFailWithUnsupportedMetricType(t *testing.T) {
// 	// arrange
// 	var (
// 		expectedResult string = ""
// 		expectedError  error  = generator.ErrUnsupportedMetricType

// 		input = "invalid metric type placeholder"
// 	)

// 	// act
// 	var result, err = generator.EmitOpenMetrics(input)

// 	// assert
// 	t.Run("simple metrics: success", func(t *testing.T) {
// 		assert.Equal(t,
// 			expectedResult,
// 			result,
// 		)

// 		assert.Equal(t,
// 			expectedError,
// 			err,
// 		)
// 	})
// }
