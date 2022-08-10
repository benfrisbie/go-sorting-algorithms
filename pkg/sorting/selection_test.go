package sorting

import (
	"sort"
	"testing"
)

func TestSelectionInt(t *testing.T) {
	s := RandomIntSlice(SliceLength)
	SelectionSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("SelectionSort failed - %v", s)
	}
}

func TestSelectionFloat64(t *testing.T) {
	s := RandomFloat64Slice(SliceLength)
	SelectionSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("SelectionSort failed - %v", s)
	}
}

func TestSelectionString(t *testing.T) {
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
