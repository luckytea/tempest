package generator_test

import (
	"testing"
	"time"

	"github.com/luckytea/tempest/generator"

	"github.com/stretchr/testify/require"
)

// Benchmark_GenerateSamplesFromString-12    	  137696	      8440 ns/op	     408 B/op	       5 allocs/op #"code,200,3".
func Benchmark_GenerateSamplesFromString(b *testing.B) {
	var (
		labels = "code,200,3;code,404,5"

		from int64 = 1620388800
		to   int64 = 1620388900

		err error
	)

	genProvider := generator.New(time.Now().Unix())

	for n := 0; n < b.N; n++ {
		_, err = genProvider.GenerateSamplesFromString(labels, from, to)
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
				{LabelName: "code", LabelValue: "200", Value: 1, Timestamp: 1620388805},
				{LabelName: "code", LabelValue: "200", Value: 2, Timestamp: 1620388827},
				{LabelName: "code", LabelValue: "200", Value: 3, Timestamp: 1620388852},
			}
		)

		genProvider := generator.New(0)

		// act
		got, err := genProvider.GenerateSamplesFromString(labels, from, to)

		// assert
		require.NoError(t, err)
		require.Equal(t, want, got)
	})
}

func Test_GenerateSamplesFromString_Success_Multiple_labels(t *testing.T) {
	t.Run("GenerateSamplesFromString: success with multiple labels", func(t *testing.T) {
		// arrange
		var (
			labels       = "code,200,3;code,404,5"
			from   int64 = 1620388800
			to     int64 = 1620388900

			want = []generator.BackfillSample{
				{LabelName: "code", LabelValue: "404", Value: 1, Timestamp: 1620388802},
				{LabelName: "code", LabelValue: "200", Value: 1, Timestamp: 1620388805},
				{LabelName: "code", LabelValue: "200", Value: 2, Timestamp: 1620388827},
				{LabelName: "code", LabelValue: "200", Value: 3, Timestamp: 1620388852},
				{LabelName: "code", LabelValue: "404", Value: 2, Timestamp: 1620388853},
				{LabelName: "code", LabelValue: "404", Value: 3, Timestamp: 1620388856},
				{LabelName: "code", LabelValue: "404", Value: 4, Timestamp: 1620388863},
				{LabelName: "code", LabelValue: "404", Value: 5, Timestamp: 1620388894},
			}
		)

		genProvider := generator.New(0)

		// act
		got, err := genProvider.GenerateSamplesFromString(labels, from, to)

		// assert
		require.NoError(t, err)
		require.Equal(t, want, got)
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

		genProvider := generator.New(time.Now().Unix())

		// act
		got, err := genProvider.GenerateSamplesFromString(labels, from, to)

		// assert
		require.ErrorIs(t,
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

		genProvider := generator.New(time.Now().Unix())

		// act
		got, err := genProvider.GenerateSamplesFromString(labels, from, to)

		// assert
		require.ErrorIs(t,
			wantError, err)

		require.Equal(t,
			want, got)
	})
}
