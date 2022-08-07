package sorting

import (
	"golang.org/x/exp/constraints"
)

func BubbleSort[T constraints.Ordered](d []T) {
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d)-i-1; j++ {
			if d[j] > d[j+1] {
				d[j], d[j+1] = d[j+1], d[j]
			}
		}
	}
}
