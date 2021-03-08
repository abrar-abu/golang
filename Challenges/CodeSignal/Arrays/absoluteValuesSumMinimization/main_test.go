package main

import (
	"fmt"
	"testing"
)

func TestAbsoluteValuesSumMinimizationSingel(t *testing.T) {
	ans := absoluteValuesSumMinimization([]int{2, 4, 7})
	if ans != 4 {
		t.Errorf("absoluteValuesSumMinimization([]int{2, 4, 7}) = %v; want 4", ans)
	}
}

func TestAbsoluteValuesSumMinimizationTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
	}{
		{[]int{2, 4, 7}, 4},
		{[]int{2, 3}, 2},
		{[]int{1, 1, 3, 4}, 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := absoluteValuesSumMinimization(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
