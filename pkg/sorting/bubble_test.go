package sorting

import (
	"sort"
	"testing"
)

func TestBubbleInt(t *testing.T) {
	s := RandomIntSlice(SliceLength)
	BubbleSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("BubbleSort failed - %v", s)
	}
}

func TestBubbleFloat64(t *testing.T) {
	s := RandomFloat64Slice(SliceLength)
	BubbleSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("BubbleSort failed - %v", s)
	}
}

func TestBubbleString(t *testing.T) {
	s := RandomStringSlice(SliceLength)
	BubbleSort(s)
	if !sort.IsSorted(sort.StringSlice(s)) {
		t.Errorf("BubbleSort failed - %v", s)
	}
}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := RandomIntSlice(SliceLength)
		BubbleSort(s)
	}
}
