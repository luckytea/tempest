// Package generator contains methods for generating metrics.
package generator

import (
	"math/rand"
)

func New(seed int64) *Provider {
	return &Provider{
		RandSource: rand.New(rand.NewSource(seed)), // nolint: gosec // not used for security purposes.
	}
}
