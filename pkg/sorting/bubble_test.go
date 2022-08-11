package sorting

import (
	"fmt"
	"sort"
	"testing"
)

func ExampleBubbleSort() {
	s := []int{5, 4, 3, 2, 1}
	BubbleSort(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5]
}

func TestBubbleSortInt(t *testing.T) {
	s := RandomIntSlice(SliceLength)
	BubbleSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("BubbleSort failed - %v", s)
	}
}

func TestBubbleSortFloat64(t *testing.T) {
	s := RandomFloat64Slice(SliceLength)
	BubbleSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("BubbleSort failed - %v", s)
	}
}

func TestBubbleSortString(t *testing.T) {
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
