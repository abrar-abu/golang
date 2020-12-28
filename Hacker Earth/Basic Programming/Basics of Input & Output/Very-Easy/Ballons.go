/*
You are conducting a contest at your college. This contest consists of two problems and  participants. You know the problem that a candidate will solve during the contest.

You provide a balloon to a participant after he or she solves a problem. There are only green and purple-colored balloons available in a market. Each problem must have a balloon associated with it as a prize for solving that specific problem. You can distribute balloons to each participant by performing the following operation:

Use green-colored balloons for the first problem and purple-colored balloons for the second problem
Use purple-colored balloons for the first problem and green-colored balloons for the second problem
You are given the cost of each balloon and problems that each participant solve. Your task is to print the minimum price that you have to pay while purchasing balloons.

Input format

First line: T that denotes the number of test cases (1<=T<=10)
For each test case: 
First line: Cost of green and purple-colored balloons 
Second line: n that denotes the number of participants (1<=n<=10)
Next n lines: Contain the status of users. For example, if the value of the j integer in the i row is 0, then it depicts that the i participant has not solved the j problem. Similarly, if the value of the j integer in the i row is 1, then it depicts that the i participant has solved the j problem.
Output format
For each test case, print the minimum cost that you have to pay to purchase balloons.

SAMPLE INPUT 
2
9 6
10
1 1
1 1
0 1
0 0
0 1
0 0
0 1
0 1
1 1
0 0
1 9
10
0 1
0 0
0 0
0 1
1 0
0 1
0 1
0 0
0 1
0 0

SAMPLE OUTPUT 
69
14

*/

package main 

import (
	"fmt"
)

func main() {

	var testCase int 
	fmt.Scan(&testCase)
	// fmt.Println("Test case are ",testCase)
	for i:=1;i<=testCase && i>=1 && i<=10;i++{
		var greenBaloonCost, purpleBallonCost int 
		fmt.Scan(&greenBaloonCost, &purpleBallonCost)
		// fmt.Println("Green and Purple ballons cost are ",greenBaloonCost,purpleBallonCost)
		var participantCount int
		fmt.Scan(&participantCount)
		// fmt.Println("Participants are ",participantCount)
		// status := make([][2]int)
		var countFirst int
		var countSecond int
		for j:=1; i<=participantCount && j<=10 && j>=1 ; j++ {
			var first int
			fmt.Scan(&first)
			// fmt.Println("First value is ",first)
			if(first == 1){
				countFirst++
			}	
			var second int
			fmt.Scan(&second)
			// fmt.Println("Second value is ",second)
			if(second == 1){
				countSecond++
			}	
		}
		min1 := countFirst * greenBaloonCost + countSecond * purpleBallonCost;
 		min2 := countFirst * purpleBallonCost + countSecond * greenBaloonCost;
		// fmt.Println("min1 and min2",min1,min2)
		// fmt.Println("countFirst and countSecond",countFirst,countSecond)
		if(min1 < min2){
			fmt.Println(min1)
		} else {
			fmt.Println(min2)
		}
		// fmt.Println("min1 and min2",min1,min2)
	}
}