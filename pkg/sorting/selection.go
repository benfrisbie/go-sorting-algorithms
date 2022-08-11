package sorting

import (
	"golang.org/x/exp/constraints"
)

// SelectionSort performs an in-place selection sort on the slice d
func SelectionSort[T constraints.Ordered](d []T) {
	for i := 0; i < len(d); i++ {
		minIndex := i
		for j := i + 1; j < len(d); j++ {
			if d[minIndex] > d[j] {
				minIndex = j
			}
		}
		d[i], d[minIndex] = d[minIndex], d[i]
	}
}
