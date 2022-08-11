package sorting

import (
	"fmt"
	"sort"
	"testing"
)

func ExampleInsertionSort() {
	s := []int{5, 4, 3, 2, 1}
	InsertionSort(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5]
}

func TestInsertionSortInt(t *testing.T) {
	s := RandomIntSlice(SliceLength)
	InsertionSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("InsertionSort failed - %v", s)
	}
}

func TestInsertionSortFloat64(t *testing.T) {
	s := RandomFloat64Slice(SliceLength)
	InsertionSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("InsertionSort failed - %v", s)
	}
}

func TestInsertionSortString(t *testing.T) {
	s := RandomStringSlice(SliceLength)
	InsertionSort(s)
	if !sort.IsSorted(sort.StringSlice(s)) {
		t.Errorf("InsertionSort failed - %v", s)
	}
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := RandomIntSlice(SliceLength)
		InsertionSort(s)
	}
}
