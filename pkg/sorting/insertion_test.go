package sorting

import (
	"sort"
	"testing"
)

func TestInsertion(t *testing.T) {
	d := []int{5, 4, 3, 2, 1}
	InsertionSort(d)
	if !sort.IsSorted(sort.IntSlice(d)) {
		t.Errorf("InsertionSort failed - %v", d)
	}
}
