package main

import "fmt"

func MinRotations(array []int) int {
	count := 0
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			fmt.Println(array[i], array[i+1])
			count = i + 1
		}
	}
	return count
}

func main() {
	fmt.Println("Min Rotation Challenge")

	testArr := []int{15, 18, 2, 3, 6, 12}
	min := MinRotations(testArr) // returns 2
	fmt.Println(min)

	testArr2 := []int{7, 9, 11, 12, 5}
	min2 := MinRotations(testArr2) // return 4
	fmt.Println(min2)

}
