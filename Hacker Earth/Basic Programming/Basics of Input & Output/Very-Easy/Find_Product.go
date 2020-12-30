/*
You have been given an array A of size N consisting of positive integers. You need to find and print the product of all the number in this array Modulo .

Input Format:
The first line contains a single integer N denoting the size of the array. The next line contains N space separated integers denoting the elements of the array

Output Format:
Print a single integer denoting the product of all the elements of the array Modulo 10^9 +7.

Constraints:
1 <= N 		<=10^3
1 <= A[i] 	<=10^3

SAMPLE INPUT 	
5
1 2 3 4 5

SAMPLE OUTPUT 	
120


*/

// Write your code here
package main

import (
	"fmt"
)

func main () {
	var answer int64 = 1
	var i int64 = 1
	var number int64
	var tempInput int64
	// array := make([]int64, 1)
	fmt.Scan(&number)
	for i <=number {
		fmt.Scan(&tempInput)
		// array = append(array,tempInput)
		answer = answer * tempInput %(1000000007)
		i++
	}
	fmt.Println(answer)

}