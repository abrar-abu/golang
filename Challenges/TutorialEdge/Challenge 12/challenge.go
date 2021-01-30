package main

import (
	"fmt"
	"math"
)

type MyInt int

func (n *MyInt) IsArmstrong() bool {
	var check MyInt
	var number MyInt
	number = *n
	length := number.length()

	for number > 0 {
		c := number % 10
		number /= 10
		check += MyInt(math.Pow(float64(c), float64(length)))
	}
	if *n == check {
		return true
	}
	return false
}

func (n *MyInt) length() (length int) {
	var number MyInt = *n
	length = 0
	for number > 0 {
		number /= 10
		length++
	}
	return
}


func main() {
  	fmt.Println("Armstrong Numbers")

  	var num1 MyInt = 371
	fmt.Println(num1.IsArmstrong())

	/*  
	var i MyInt = 1
	for  i < 10000000 {
		if(i.IsArmstrong()){
			fmt.Println(i)
		}
		i = i +1
	}
	*/
}
