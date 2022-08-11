package sorting

import (
	"fmt"
	"sort"
	"testing"
)

// ExampleBubbleSort calls BubbleSort on the slice [5, 4, 3, 2, 1] and expects output of [1 2 3 4 5]
func ExampleBubbleSort() {
	s := []int{5, 4, 3, 2, 1}
	BubbleSort(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5]
}

// TestBubbleSortInt calls BubbleSort on a slice of random ints
func TestBubbleSortInt(t *testing.T) {
	s := RandomIntSlice(*SliceLength)
	t.Log(*SliceLength)
	BubbleSort(s)
	if !sort.IsSorted(sort.IntSlice(s)) {
		t.Errorf("BubbleSort failed - %v", s)
	}
}

// TestBubbleSortFloat64 calls BubbleSort on a slice of random float64s
func TestBubbleSortFloat64(t *testing.T) {
	s := RandomFloat64Slice(*SliceLength)
	BubbleSort(s)
	if !sort.IsSorted(sort.Float64Slice(s)) {
		t.Errorf("BubbleSort failed - %v", s)
	}
}

// TestBubbleSortString calls BubbleSort on a slice of random strings
func TestBubbleSortString(t *testing.T) {
	s := RandomStringSlice(*SliceLength)
	BubbleSort(s)
	if !sort.IsSorted(sort.StringSlice(s)) {
		t.Errorf("BubbleSort failed - %v", s)
	}
}

// BenchmarkBubble calls BubbleSort on a slice of random ints repeatedly until a consistent benchmark is reached
func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := RandomIntSlice(*SliceLength)
		BubbleSort(s)
	}
}
