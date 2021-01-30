package main

import "fmt"

func TriangularNumber(n int) int {
  	// https://en.wikipedia.org/wiki/Triangular_number
  	return (n *(n+1))/2
}

func main() {
  	fmt.Println("Returning the 'nth' triangular number")

  	number := TriangularNumber(3)
  	fmt.Println(number) // '6'
	for  index := 1; index <= 15; index++  {
		fmt.Println(TriangularNumber(index))
	}
}
