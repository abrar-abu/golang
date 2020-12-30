/*
You have been given a String S. You need to find and print whether this string is a palindrome or not. If yes, print "YES" (without quotes), else print "NO" (without quotes).

Input Format
The first and only line of input contains the String S. The String shall consist of lowercase English alphabets only.

Output Format
Print the required answer on a single line.

Constraints 
1 <= S <=10

SAMPLE INPUT 	
aba
SAMPLE OUTPUT 	
YES

Note
String S consists of lowercase English Alphabets only.

*/
package main

import (
   "fmt"
   "strings"
)
func Reverse(s string) (result string) {
   for _, v := range s {
      result = string(v) + result
   }
   return
}

func IsPalindrome(str string) interface{} {
   if str == Reverse(str) {
      return true
   }
   return false
}
func main() {
   var str string
   // fmt.Print("Enter a string: ")
   fmt.Scan(&str)
   if IsPalindrome(strings.ToUpper(str)) == true {
      fmt.Print("YES")
   } else {
      fmt.Print("NO")
   }
}