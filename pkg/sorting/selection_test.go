package sorting

import (
	"fmt"
	"sort"
	"testing"
)

func ExampleSelectionSort() {
	s := []int{5, 4, 3, 2, 1}
	SelectionSort(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5]
}

func TestSelectionSortInt(t *testing.T) {
	s := RandomIntSlice(SliceLength)
	SelectionSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("SelectionSort failed - %v", s)
	}
}

func TestSelectionSortFloat64(t *testing.T) {
	s := RandomFloat64Slice(SliceLength)
	SelectionSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("SelectionSort failed - %v", s)
	}
}

func TestSelectionSortString(t *testing.T) {
	s := RandomStringSlice(SliceLength)
	SelectionSort(s)
	if !sort.IsSorted(sort.StringSlice(s)) {
		t.Errorf("SelectionSort failed - %v", s)
	}
}

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := RandomIntSlice(SliceLength)
		SelectionSort(s)
	}
}
