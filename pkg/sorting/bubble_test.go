package sorting

import (
	"fmt"
	"sort"
	"testing"
)

func TestBubble(t *testing.T) {
	tests := [][]int{
		{5, 4, 3, 2, 1},
		{1, 2, 5, 4, 3},
		{-1, -2, 5, 4, 3},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test), testIntBubble(test))
	}
}

func testIntBubble(d []int) func(*testing.T) {
	return func(t *testing.T) {
		BubbleSort(d)
		if !sort.IsSorted(sort.IntSlice(d)) {
			t.Errorf("BubbleSort failed - %v", d)
		}
	}
}
