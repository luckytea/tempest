package cfg_test

import (
	"testing"

	"github.com/LuckyTea/tempest/cfg"
	"github.com/LuckyTea/tempest/model"
	"github.com/stretchr/testify/require"
)

func TestConfig_Validate_Successful(t *testing.T) {
	t.Run("Config.Validate: Successful", func(t *testing.T) {
		// arrange
		cfg := &cfg.Config{
			Name: "test_metric",
			Type: "counter",
		}

		// act
		_, err := cfg.Validate()

		// assert
		require.NoError(t, err)
	})
}

func TestConfig_Validate_UnsupportedMetricType(t *testing.T) {
	t.Run("Config.Validate: Unsupported Metric Type", func(t *testing.T) {
		// arrange
		cfg := &cfg.Config{
			Name: "test_metric",
			Type: "gauge",
		}

		// act
		_, err := cfg.Validate()

		// assert
		require.Equal(t,
			model.ErrUnsupportedMetricType,
			err,
		)
	})
}
