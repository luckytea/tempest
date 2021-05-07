package generator

import (
	"sort"
	"strconv"
	"strings"
)

func (p *Provider) GenerateSamplesFromString(labelsString string, from, to int64) ([]BackfillSample, error) {
	var samples []BackfillSample

	labels := strings.Split(labelsString, ";")

	for _, label := range labels {
		tmp := strings.Split(label, ",")

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
				Timestamp:  p.RandSource.Int63n(to-from) + from,
			}

			samples = append(samples, bs)
		}
	}

	sort.Sort(TimestampSorter(samples))

	m := make(map[string]float64)

	for ii := range samples {
		m[samples[ii].LabelValue]++

		samples[ii].Value = m[samples[ii].LabelValue]
	}

	return samples, nil
}

type TimestampSorter []BackfillSample

func (a TimestampSorter) Len() int           { return len(a) }
func (a TimestampSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TimestampSorter) Less(i, j int) bool { return a[i].Timestamp < a[j].Timestamp }
