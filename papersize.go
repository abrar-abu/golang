package main 

import "fmt"

func main() {

	var length int64  = 1189
	var breadth int64 = 841

	for  i:=0; i<=8; i++ {
		if length>breadth {
			fmt.Println()
			length = length/2;
		} else {
			breadth = breadth/2;
		}
		fmt.Println(length,"x",breadth)
	}
}