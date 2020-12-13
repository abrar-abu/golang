/*
Matthieu is an expert in Mathematics.He is performing arithmetic operations and wants to find all the divisors of P number.
Note : Divisors should be separated by space and increasing order should be maintained .
Input Format
Input contains a non negative integer "P" denoting the number
	Violation of the Input criteria: System should display message as "Wrong input".
Output Format
All divisors should be separated by space and increasing order must be maintained.

Constraints
1<=P<=10^8

Sample 1:
Input
10
Output
1 2 5 10

Sample 2:
Input
18
Output
1 2 3 6 9 18
*/
package main 

import "fmt"

func divisor(num int) {
	for index := 1; index <= num; index++ {
		// fmt.Println(num%index)
		if num%index == 0 {
			fmt.Println(index)
		}
	}
}

func divisors(num int) ([]int) {
	// var divisors [] int
	var divisors = make([]int,num/2)
	iterator := 0
	for index := 1; index <= num; index++ {
		// fmt.Println(num%index)
		if num%index == 0 {
			divisors[iterator] = index
			iterator++
		}
	}
	return divisors
}

func main() {
	fmt.Println("Test the data ")
	divisor(10)
	fmt.Println(divisors(12))
}