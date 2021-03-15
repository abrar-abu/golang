package main

import "fmt"

func arrayChange(inputArray []int) int {
	var counter int
	var temp = inputArray[0]
	for index := 1; index < len(inputArray); index++ {
		if inputArray[index] <= temp {
			counter = counter + (temp - inputArray[index]) + 1
			//fmt.Println(counter)
			temp = temp + 1
			//inputArray[index] = temp
		} else {
			temp = inputArray[index]
		}
		//fmt.Println(inputArray)
	}

	// fmt.Println(inputArray)
	// fmt.Println(counter)
	return counter
}

func main() {
	fmt.Printf("%v", arrayChange([]int{1, 2, 3}))
}
