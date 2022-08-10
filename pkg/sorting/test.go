package sorting

import (
	"math/rand"
)

var SliceLength = 1000

func RandomIntSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Int()
	}
	return s
}

func RandomFloat64Slice(n int) []float64 {
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Float64()
	}
	return s
}

func RandomStringSlice(n int) []string {
	var runes = []rune("abcdefghijklmnopqrstuvwxyz")
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = randomString(rand.Intn(len(runes)), runes)
	}
	return s
}

func randomString(n int, runes []rune) string {

	s := make([]rune, n)
	for i := range s {
		s[i] = runes[rand.Intn(len(runes))]
	}
	return string(s)
}
