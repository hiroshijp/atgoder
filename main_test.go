package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{-2, -3, -2, -5, -1, -5}, 1},
		{[]int{1, 3, 6, 4, 1, 2}, 5},
		{[]int{1, 2, 3}, 4},
		{[]int{-1, -3}, 1},
		{[]int{2, 3, 7, 6, 8, -1, -10, 15}, 1},
		{[]int{1, 2, 0}, 3},
	}

	for _, test := range tests {
		result := Solution(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	input := []int{2, 3, 7, 6, 8, -1, -10, 15}
	for i := 0; i < b.N; i++ {
		Solution(input)
	}
}
