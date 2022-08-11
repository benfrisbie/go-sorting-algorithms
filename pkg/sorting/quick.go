package sorting

import (
	"golang.org/x/exp/constraints"
)

// QuickSort performs an in-place quick sort on the slice d
func QuickSort[T constraints.Ordered](d []T) {
	if len(d) <= 1 {
		return
	}

	// Partition
	pivotIndex := 0
	pivotValue := d[len(d)-1]
	i := 0
	for ; i < len(d); i++ {
		if d[i] < pivotValue {
			d[pivotIndex], d[i] = d[i], d[pivotIndex]
			pivotIndex++
		}
	}
	d[pivotIndex], d[len(d)-1] = d[len(d)-1], d[pivotIndex]

	// QuickSort to the left and right of the pivot
	QuickSort(d[:pivotIndex])
	QuickSort(d[pivotIndex+1:])
}
