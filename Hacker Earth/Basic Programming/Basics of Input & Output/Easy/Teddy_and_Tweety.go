/*
Teddy has N chocolates. Tweety asks him to give few chocolates to her. Teddy is very generous and kind hearted. So he decides to give all chocolates to Tweety, but only on one condition.

If integer N can be divided into 3 parts, such that these three parts can form sides of an equilateral triangle, then Teddy gives all chocolates to Tweety. Otherwise he won't give any chocolate to her.

Print "YES" (without quotes) if Teddy will give all chocolates to Tweety. Otherwise print "NO" (without quotes).

Input:

Integer N

Output:

YES or NO depending on problem.

Constraints:

1 <= N  <= 10000

SAMPLE INPUT 
6
SAMPLE OUTPUT 
YES
Explanation
6 can be divided in three parts - 2, 2, 2. These can form sides on equilateral triangle.
*/

package main

import (
    "fmt"
)

func main (){
    var chocolates int
    fmt.Scan(&chocolates)
    if(chocolates%3 == 0){
        fmt.Println("YES")
    }else {
        fmt.Println("NO")
    }
}