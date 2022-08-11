package sorting

import (
	"golang.org/x/exp/constraints"
)

// MergeSort performs an in-place merge sort on the slice d
func MergeSort[T constraints.Ordered](d []T) {
	if len(d) <= 1 {
		return
	}

	// Split the array in half
	mid := len(d) / 2
	MergeSort(d[:mid])
	MergeSort(d[mid:])

	// Merge the two sorted arrays
	i := 0
	j := mid
	for i < mid && j < len(d) {
		if d[i] <= d[j] {
			i++
		} else {
			v := d[j]
			for k := j; k > i; k-- {
				d[k] = d[k-1]
			}
			d[i] = v
			i++
			j++
			mid++
		}

	}
}
