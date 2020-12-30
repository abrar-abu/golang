/*
Consider All lowercase Alphabets of the English language. Here we consider each alphabet from a to z to have a certain weight. The weight of the alphabet a is considered to be 1, b to be 2, c to be 3 and so on until z has a weight of 26. In short, the weight of the alphabet a is 1, and the weight of all other alphabets is the weight of its previous alphabet + 1.

Now, you have been given a String S consisting of lowercase English characters. You need to find the summation of weight of each character in this String.

For example, Consider the String aba

Here, the first character a has a weight of 1, the second character b has 2 and the third character a again has a weight of 1. So the summation here is equal to : 1+2+1=4

Input Format:

The first and only line of input contains the String S.

Output Format:

Print the required answer on a single line

Constraints:

1≤|S|≤100

SAMPLE INPUT 
aba
SAMPLE OUTPUT 
4

*/

// Write your code here
package main

import (
    "fmt"
    "strings"
)

func main () {
    var word string;
    fmt.Scan(&word)
    wordLowerCase := strings.ToLower(word) 
    wordLowerCaseCharacter := []byte(wordLowerCase)
    var sum int
    for i,_:= range wordLowerCaseCharacter{
		// fmt.Println(value)
		sum += (int(wordLowerCaseCharacter[i]) - 96)
    }
    fmt.Println(sum)
}