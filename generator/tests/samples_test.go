package generator_test

import (
	"testing"

	"github.com/luckytea/tempest/generator"

	"github.com/stretchr/testify/require"
)

// Benchmark_GenerateSamplesFromString-12    	  137696	      8440 ns/op	     408 B/op	       5 allocs/op #"code,200,3".
func Benchmark_GenerateSamplesFromString(b *testing.B) {
	var (
		labels = "code,200,3"

		from int64 = 1620388800
		to   int64 = 1620388900

		err error
	)

	for n := 0; n < b.N; n++ {
		_, err = generator.GenerateSamplesFromString(labels, from, to)
		if err != nil {
			b.Fail()
		}
	}
}

func Test_GenerateSamplesFromString_Success(t *testing.T) {
	t.Run("GenerateSamplesFromString: success with single label", func(t *testing.T) {
		// arrange
		var (
			labels       = "code,200,3"
			from   int64 = 1620388800
			to     int64 = 1620388900

			want = []generator.BackfillSample{
				{LabelName: "code", LabelValue: "200", Value: 1},
				{LabelName: "code", LabelValue: "200", Value: 2},
				{LabelName: "code", LabelValue: "200", Value: 3},
			}
		)

		// act
		got, err := generator.GenerateSamplesFromString(labels, from, to)

		// assert
		require.NoError(t, err)

		for i := range got {
			require.Equal(t, want[i].LabelName, got[i].LabelName)
			require.Equal(t, want[i].LabelValue, got[i].LabelValue)
			require.Equal(t, want[i].Value, got[i].Value)

			require.LessOrEqual(t, got[i].Timestamp, to)
			require.GreaterOrEqual(t, got[i].Timestamp, from)
		}
	})
}

func Test_GenerateSamplesFromString_Fail_malformed(t *testing.T) {
	t.Run("GenerateSamplesFromString: fail - malformed label string", func(t *testing.T) {
		// arrange
		var (
			labels       = "code,200"
			from   int64 = 1620388800
			to     int64 = 1620388900

			want      []generator.BackfillSample = nil
			wantError                            = generator.ErrMalformed
		)

		// act
		got, err := generator.GenerateSamplesFromString(labels, from, to)

		// assert
		require.Equal(t,
			wantError, err)

		require.Equal(t,
			want, got)
	})
}

func Test_GenerateSamplesFromString_Fail_malformed_count(t *testing.T) {
	t.Run("GenerateSamplesFromString: fail - malformed count in string", func(t *testing.T) {
		// arrange
		var (
			labels       = "code,200,a"
			from   int64 = 1620388800
			to     int64 = 1620388900

			want      []generator.BackfillSample = nil
			wantError                            = generator.ErrMalformed
		)

		// act
		got, err := generator.GenerateSamplesFromString(labels, from, to)

		// assert
		require.Equal(t,
			wantError, err)

		require.Equal(t,
			want, got)
	})
}
