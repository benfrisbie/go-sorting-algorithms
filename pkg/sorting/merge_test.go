package sorting

import (
	"fmt"
	"sort"
	"testing"
)

// ExampleMergeSort calls MergeSort on the slice [5, 4, 3, 2, 1] and expects output of [1 2 3 4 5]
func ExampleMergeSort() {
	s := []int{5, 4, 3, 2, 1}
	MergeSort(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5]
}

// TestMergeSortInt calls MergeSort on a slice of random ints
func TestMergeSortInt(t *testing.T) {
	s := RandomIntSlice(*SliceLength)
	MergeSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("MergeSort failed - %v", s)
	}
}

// TestMergeSortFloat64 calls MergeSort on a slice of random float64s
func TestMergeSortFloat64(t *testing.T) {
	s := RandomFloat64Slice(*SliceLength)
	MergeSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("MergeSort failed - %v", s)
	}
}

// TestMergeSortString calls MergeSort on a slice of random strings
func TestMergeSortString(t *testing.T) {
	s := RandomStringSlice(*SliceLength)
	MergeSort(s)
	if !sort.IsSorted(sort.StringSlice(s)) {
		t.Errorf("MergeSort failed - %v", s)
	}
}

// BenchmarkMerge calls MergeSort on a slice of random ints repeatedly until a consistent benchmark is reached
func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := RandomIntSlice(*SliceLength)
		MergeSort(s)
	}
}
