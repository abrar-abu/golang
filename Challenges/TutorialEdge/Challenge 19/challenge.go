package main

import "fmt"

func DiffSquares(n, m int) int {
  return (n*n) - (m*m)
}

func main() {
  fmt.Println("Calculate The Difference of Squares",DiffSquares(25,13))
}
