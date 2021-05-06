package generator

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GenerateSamplesFromString(labels string, from, to int64) ([]BackfillSample, error) {
	var samples []BackfillSample

	rand.Seed(time.Now().Unix())

	tmp := strings.Split(labels, ",")

	if len(tmp)%3 != 0 {
		return nil, ErrMalformed
	}

	count, err := strconv.Atoi(tmp[2])
	if err != nil {
		return nil, ErrMalformed
	}

	for ii := 0; ii < count; ii++ {
		var bs = BackfillSample{
			LabelName:  tmp[0],
			LabelValue: tmp[1],
			Timestamp:  rand.Int63n(to-from) + from, // nolint: gosec // not used for security purposes.
		}

		samples = append(samples, bs)
	}

	sort.Sort(TimestampSorter(samples))

	for ii := range samples {
		samples[ii].Value = float64(ii) + 1
	}

	return samples, nil
}

type TimestampSorter []BackfillSample

func (a TimestampSorter) Len() int           { return len(a) }
func (a TimestampSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TimestampSorter) Less(i, j int) bool { return a[i].Timestamp < a[j].Timestamp }
