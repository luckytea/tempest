package cfg_test

import (
	"testing"

	"github.com/luckytea/tempest/cfg"
	"github.com/luckytea/tempest/model"

	"github.com/stretchr/testify/require"
)

func TestConfig_Validate_Successful(t *testing.T) {
	t.Run("Config.Validate: Successful", func(t *testing.T) {
		// arrange
		config := &cfg.Config{
			Type:  "counter",
			Name:  "test_metric",
			Desc:  "test_desc",
			Label: "name,value,3",
			From:  500,
			To:    600,
		}

		// act
		_, err := config.Validate()

		// assert
		require.NoError(t, err)
	})
}

func TestConfig_Validate_UnsupportedMetricType(t *testing.T) {
	t.Run("Config.Validate: Unsupported Metric Type", func(t *testing.T) {
		// arrange
		config := &cfg.Config{
			Name: "test_metric",
			Type: "gauge",
		}

		// act
		_, err := config.Validate()

		// assert
		require.Equal(t,
			model.ErrUnsupportedMetricType,
			err,
		)
	})
}

func TestConfig_Validate_MalformedTime(t *testing.T) {
	t.Run("Config.Validate: malformed time configuration", func(t *testing.T) {
		// arrange
		config := &cfg.Config{
			Type: "counter",
			Name: "test_metric",
			From: 600,
			To:   500,
		}

		// act
		_, err := config.Validate()

		// assert
		require.Equal(t,
			cfg.ErrMalformedTime,
			err,
		)
	})
}
