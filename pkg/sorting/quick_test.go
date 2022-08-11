package sorting

import (
	"fmt"
	"sort"
	"testing"
)

// ExampleQuickSort calls QuickSort on the slice [5, 4, 3, 2, 1] and expects output of [1 2 3 4 5]
func ExampleQuickSort() {
	s := []int{5, 4, 3, 2, 1}
	QuickSort(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5]
}

// TestQuickSortInt calls QuickSort on a slice of random ints
func TestQuickSortInt(t *testing.T) {
	s := RandomIntSlice(*SliceLength)
	QuickSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("QuickSort failed - %v", s)
	}
}

// TestQuickSortFloat64 calls QuickSort on a slice of random float64s
func TestQuickSortFloat64(t *testing.T) {
	s := RandomFloat64Slice(*SliceLength)
	QuickSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("QuickSort failed - %v", s)
	}
}

// TestQuickSortString calls QuickSort on a slice of random strings
func TestQuickSortString(t *testing.T) {
	s := RandomStringSlice(*SliceLength)
	QuickSort(s)
	if !sort.IsSorted(sort.StringSlice(s)) {
		t.Errorf("QuickSort failed - %v", s)
	}
}

// BenchmarkQuick calls QuickSort on a slice of random ints repeatedly until a consistent benchmark is reached
func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := RandomIntSlice(*SliceLength)
		QuickSort(s)
	}
}
