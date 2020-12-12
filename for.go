package main

import "fmt"

func main() {

	// 1nitialization
	i := 1
	// condition
    for i <= 3 {
		fmt.Println(i)
		// updation/postcondition
        i = i + 1
    }

	// normal for loop
	// declaration/1nitialization,condition,updation/postcondition
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

	// infinite for loop we have break
    for {
        fmt.Println("loop")
        break
    }

	// loop to print odd numbers
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
	}

	//The range-for loop
	// with index
	var items []int = []int{1, 2, 3, 4, 5}
    for index, value := range items {
        fmt.Println("[",index,"]", value)
    }

	// without index
    for _, v := range items {
        fmt.Println(v)
	}
	
	// To print the map (key value)
	var m map[int]string = map[int]string{
        1: "One",
        2: "Two",
        3: "Three",
    }
     
    for key, value := range m {
		fmt.Println("[",key,"]", value)
    }
}