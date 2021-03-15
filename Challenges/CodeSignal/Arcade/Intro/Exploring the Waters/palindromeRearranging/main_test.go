package main

import (
	"fmt"
	"testing"
)

func Equal(a, b []int) bool {
	// fmt.Printf("%T %T", a, b)
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		// fmt.Printf("v => %v b => %v\n", v, b[i])
		//fmt.Printf("v(Type) => %T b(Type) => %T\n", v, b[i])
		// fmt.Printf("v != b[i] => %v \n", v != b[i])
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestPalindromeRearrangingSingel(t *testing.T) {
	ans := palindromeRearranging("aabb")
	if ans != true {
		t.Errorf("palindromeRearranging(\"aabb\") = %v; want true", ans)
	}
}

func TestPalindromeRearrangingTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"aabb", true},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc", false},
		{"abbcabb", true},
		{"zyyzzzzz", true},
		{"z", true},
		{"zaa", true},
		{"abca", false},
		{"abcad", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbccccaaaaaaaaaaaaa", false},
		{"abdhuierf", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := palindromeRearranging(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
