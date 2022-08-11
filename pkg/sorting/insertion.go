package sorting

import (
	"golang.org/x/exp/constraints"
)

// InsertionSort performs an in-place insertion sort on the slice d
func InsertionSort[T constraints.Ordered](d []T) {
	for i := 1; i < len(d); i++ {
		for j := i; j > 0 && d[j-1] > (d[j]); j-- {
			d[j], d[j-1] = d[j-1], d[j]
		}
	}
}
