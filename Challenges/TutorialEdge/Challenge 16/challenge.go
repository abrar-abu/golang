package main

import "fmt"

func OddEvenFactors(num int) string {
	if divisorsCount(num) % 2 == 0 {
	  return "even"
	}
	return "odd"
}
func divisors(num int) ([]int) {
	// var divisors [] int
	var divisors = make([]int,0)
	iterator := 0
	for index := 2; index <= num; index++ {
		// fmt.Println(num%index)
		if num%index == 0 {
			divisors = append(divisors,index)
			iterator++
		}
	}
	// fmt.Println(divisors)
	return divisors
}

func divisorsCount(divisor int) int {
	return len(divisors(divisor))
}


func main() {
  fmt.Println("Odd or Even Factors")

  numFactors := OddEvenFactors(23)
  fmt.Println(numFactors) // "even"

  numFactors = OddEvenFactors(36)
  fmt.Println(numFactors) // "odd"
}
