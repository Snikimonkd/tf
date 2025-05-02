package options

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/exp/constraints"
)

// SortSlicesOpt - option that sorts slices
func SortSlicesOpt[T constraints.Ordered]() cmp.Option {
	return cmpopts.SortSlices(func(x, y T) bool {
		return x < y
	})
}

// SortMapsOpt - option that sorts maps
func SortMapsOpt[T constraints.Ordered]() cmp.Option {
	return cmpopts.SortMaps(func(x, y T) bool {
		return x < y
	})
}
