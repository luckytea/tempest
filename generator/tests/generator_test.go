package generator_test

import (
	"testing"

	"github.com/luckytea/tempest/generator"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

// Test_simple_metrics_built_successfuly
// генерируем простую метрику в формате промифиуса и валидируем в формате open metrics.
func Test_simple_metrics_built_successfuly(t *testing.T) {
	// arrange
	var (
		expectedResult string = "success"
		expectedError  error  = nil

		dummy prometheus.CounterVec
	)

	// act
	var result, err = generator.EmitOpenMetrics(dummy)

	// assert
	t.Run("simple metrics: success", func(t *testing.T) {
		assert.Equal(t,
			expectedResult,
			result,
		)

		assert.Equal(t,
			expectedError,
			err,
		)
	})
}

// Test_simple_metrics_built_error_unsupported_type
// передан неподдерживаемый формат метрики.
func Test_openmetrics_generated_invalid(t *testing.T) {
	// arrange
	var (
		expectedResult string = ""
		expectedError  error  = generator.ErrUnsupportedMetricType

		dummy = "invalid metric type placeholder"
	)

	// act
	var result, err = generator.EmitOpenMetrics(dummy)

	// assert
	t.Run("simple metrics: success", func(t *testing.T) {
		assert.Equal(t,
			expectedResult,
			result,
		)

		assert.Equal(t,
			expectedError,
			err,
		)
	})
}
