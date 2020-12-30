/*
You have been given 3 integers - l, r and k. Find how many numbers between l and r (both inclusive) are divisible by k. You do not need to print these numbers, you just have to find their count.

Input Format
The first and only line of input contains 3 space separated integers l, r and k.

Output Format
Print the required answer on a single line.

Constraints
1<=l<=r<=1000
1<=k<=1000 

SAMPLE INPUT 
1 10 1
SAMPLE OUTPUT 
10

*/

package main

import (
	"fmt"
	//"log"
)

func main () {
	// var l, r, k, c int
	var i, j, k, c int  
	if    _, err := fmt.Scan(&i, &j, &k);    err != nil {
		// log.Print("  Scan for i, j & k failed, due to ", err)
		return
	}
	/*
	fmt.Scan(&l, &r, &k)
	fmt.Scan(&l)
	fmt.Scan(&r)
	fmt.Scan(&k)
	*/
	for l := i ; l<=j ; l++{
		if(l%k==0) {
			c++
		}
	}
	fmt.Println(c)
	// */

}