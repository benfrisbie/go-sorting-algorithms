package sorting

import (
	"fmt"
	"sort"
	"testing"
)

// ExampleSelectionSort calls SelectionSort on the slice [5, 4, 3, 2, 1] and expects output of [1 2 3 4 5]
func ExampleSelectionSort() {
	s := []int{5, 4, 3, 2, 1}
	SelectionSort(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5]
}

// TestSelectionSortInt calls SelectionSort on a slice of random ints
func TestSelectionSortInt(t *testing.T) {
	s := RandomIntSlice(*SliceLength)
	SelectionSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("SelectionSort failed - %v", s)
	}
}

// TestSelectionSortFloat64 calls SelectionSort on a slice of random float64s
func TestSelectionSortFloat64(t *testing.T) {
	s := RandomFloat64Slice(*SliceLength)
	SelectionSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("SelectionSort failed - %v", s)
	}
}

// TestSelectionSortString calls SelectionSort on a slice of random strings
func TestSelectionSortString(t *testing.T) {
	s := RandomStringSlice(*SliceLength)
	SelectionSort(s)
	if !sort.IsSorted(sort.StringSlice(s)) {
		t.Errorf("SelectionSort failed - %v", s)
	}
}

// BenchmarkSelection calls SelectionSort on a slice of random ints repeatedly until a consistent benchmark is reached
func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := RandomIntSlice(*SliceLength)
		SelectionSort(s)
	}
}
