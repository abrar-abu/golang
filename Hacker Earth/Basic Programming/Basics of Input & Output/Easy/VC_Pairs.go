/*
Max has a string S with length N. He needs to find the number of indices i (1≤i≤N−11≤i≤N−1) such that the i-th character of this string is a consonant and the i+1th character is a vowel. However,she is busy, so she asks for your help.

Note: The letters 'a', 'e', 'i', 'o', 'u' are vowels; all other lowercase English letters are consonants.

Input

The first line of the input contains a single integer T denoting the number of test cases. The description of T test cases follows.
The first line of each test case contains a single integer N.
The second line contains a single string S with length N.
Output

For each test case, print a single line containing one integer ― the number of occurrences of a vowel immediately after a consonant

Constraints

1≤T≤1001≤T≤100
1≤N≤1001≤N≤100
SS contains only lowercase English letters
SAMPLE INPUT 
3
6
bazeci
3
abu
1
o
SAMPLE OUTPUT 
3
1
0

*/

// Write your code here

package main 

import (
    "fmt"
)

func main(){

    var testCase int
    fmt.Scan(&testCase)
    for i:=0; i<testCase ;i++ {
        var count int
        var size int
        fmt.Scan(&size)
        var word string
        fmt.Scan(&word)
		wordCharcterArray := []rune(word)
		if(len(word)==size){
            if(size == 1) {
                fmt.Println(0)
            } else {
				for i,_ := range wordCharcterArray {
					if(!(wordCharcterArray[i]=='a'||wordCharcterArray[i]=='e'|| wordCharcterArray[i]=='i'|| wordCharcterArray[i]=='o'||wordCharcterArray[i]=='u')&&(wordCharcterArray[i+1]=='a'||wordCharcterArray[i+1]=='e'|| wordCharcterArray[i+1]=='i'|| wordCharcterArray[i+1]=='o'||wordCharcterArray[i+1]=='u')){
							count++;
						}
				}
				fmt.Println(count)		
			}
		}
    }
}