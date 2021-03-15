package main

import "fmt"

// TODO :- Need to complete the 2 failing use case  

func palindromeRearranging(inputString string) bool {
	countOfCharterMap := make(map[rune]int)
	for _, value := range inputString {
		if countOfCharterMap[value] == 0 {
			countOfCharterMap[value] = 1
		} else {
			countOfCharterMap[value] = countOfCharterMap[value] + 1
		}
	}
	var counterOfOdds int
	for _, value := range countOfCharterMap {
		if value == 1 {
			counterOfOdds = counterOfOdds + 1
		}
	}
	if counterOfOdds > 1 {
		return false
	}
	return true
}

func main() {
	fmt.Printf("%v", palindromeRearranging("aabb"))
}
