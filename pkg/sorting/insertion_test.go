package sorting

import (
	"fmt"
	"sort"
	"testing"
)

// ExampleInsertionSort calls InsertionSort on the slice [5, 4, 3, 2, 1] and expects output of [1 2 3 4 5]
func ExampleInsertionSort() {
	s := []int{5, 4, 3, 2, 1}
	InsertionSort(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5]
}

// TestInsertionSortInt calls InsertionSort on a slice of random ints
func TestInsertionSortInt(t *testing.T) {
	s := RandomIntSlice(*SliceLength)
	InsertionSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("InsertionSort failed - %v", s)
	}
}

// TestInsertionSortFloat64 calls InsertionSort on a slice of random float64s
func TestInsertionSortFloat64(t *testing.T) {
	s := RandomFloat64Slice(*SliceLength)
	InsertionSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("InsertionSort failed - %v", s)
	}
}

// TestInsertionSortString calls InsertionSort on a slice of random strings
func TestInsertionSortString(t *testing.T) {
	s := RandomStringSlice(*SliceLength)
	InsertionSort(s)
	if !sort.IsSorted(sort.StringSlice(s)) {
		t.Errorf("InsertionSort failed - %v", s)
	}
}

// BenchmarkInsertion calls InsertionSort on a slice of random ints repeatedly until a consistent benchmark is reached
func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := RandomIntSlice(*SliceLength)
		InsertionSort(s)
	}
}
