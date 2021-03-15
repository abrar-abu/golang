package main

import (
	"fmt"
	"testing"
)

func Equal(a, b []string) bool {
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

func TestAddBorderSingel(t *testing.T) {
	ans := addBorder([]string{"abc", "ded"})
	if !Equal(ans, []string{"*****",
		"*abc*",
		"*ded*",
		"*****"}) {
		t.Errorf("addBorder([]string{\"abc\", \"ded\"}) = %v; want [\"*****\",\"*abc*\",\"*ded*\",\"*****\"]", ans)
	}
}

func TestAddBorderTableDriven(t *testing.T) {
	var tests = []struct {
		a    []string
		want []string
	}{
		{[]string{"abc",
			"ded"}, []string{"*****",
			"*abc*",
			"*ded*",
			"*****"}},
		{[]string{"a"}, []string{"***",
			"*a*",
			"***"}},
		{[]string{"aa",
			"**",
			"zz"}, []string{"****",
			"*aa*",
			"****",
			"*zz*",
			"****"}},
		{[]string{"abcde",
			"fghij",
			"klmno",
			"pqrst",
			"uvwxy"}, []string{"*******",
			"*abcde*",
			"*fghij*",
			"*klmno*",
			"*pqrst*",
			"*uvwxy*",
			"*******"}},
		{[]string{"wzy**"}, []string{"*******",
			"*wzy***",
			"*******"}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := addBorder(tt.a)
			if !Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
