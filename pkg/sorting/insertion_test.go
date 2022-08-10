package sorting

import (
	"sort"
	"testing"
)

func TestInsertionInt(t *testing.T) {
	s := RandomIntSlice(SliceLength)
	InsertionSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("InsertionSort failed - %v", s)
	}
}

func TestInsertionFloat64(t *testing.T) {
	s := RandomFloat64Slice(SliceLength)
	InsertionSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("InsertionSort failed - %v", s)
	}
}

func TestInsertionString(t *testing.T) {
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
