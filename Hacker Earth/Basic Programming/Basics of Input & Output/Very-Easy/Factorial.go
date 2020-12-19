/*
You have been given a positive integer N. You need to find and print the Factorial of this number. The Factorial of a positive integer N refers to the product of all number in the range from 1 to N. You can read more about the factorial of a number here.

Input Format:
The first and only line of the input contains a single integer N denoting the number whose factorial you need to find.

Output Format
Output a single line denoting the factorial of the number N.

Constraints

1 <= N <=10

SAMPLE INPUT 	2
SAMPLE OUTPUT 	2

*/
package main
import "fmt"
var factVal uint64 = 1

var i int = 1
var n int

func factorial(n int) uint64 {   
    for i:=1; i<=n; i++ {
        factVal *= uint64(i)
    }
    return factVal  
}
func main() {
    fmt.Scan(&n)   
    fmt.Print(factorial(n))
}
