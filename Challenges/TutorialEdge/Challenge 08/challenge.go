package main

import (
	"fmt"
	"sort"
	"strings"
)

func CheckPermutations(str1, str2 string) bool {
	if len(str1)!= len(str2){
		return false
	}
	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")
	sort.Strings(s1)
	sort.Strings(s2)
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
  return true
}

func main() {
  fmt.Println("Check Permutations Challenge")

  str1 := "adcme"
  str2 := "medac"

  isPermutation := CheckPermutations(str1, str2)
  fmt.Println(isPermutation)

}
