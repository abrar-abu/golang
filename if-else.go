package main

import "fmt"

func v() int {
    return 42
}

func main() {

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }


	// If statement with initialization
	if a := v(); a == 42 {
        fmt.Println("It's 42")        // It's 42
    } else {
        fmt.Println("It's not")
    }

	// if else ladder 
	var v int = 12
    if v > 0 && v < 10 {
        fmt.Println("0-10")
    } else if v > 10 && v < 20 {
        fmt.Println("10-20")
    } else if  v > 20 && v < 30 {
        fmt.Println("20-30")
    } else {
        fmt.Println("30+")
	}    
	
	// Nested if else 
	var name string = "Rob"
    var age int     = 63
     
    if name == "Rob" {
        if age == 64 {
            fmt.Println("He is Rob.")
        } else {
            fmt.Println("Wrong person")
        }
    } else {
        fmt.Println("Who are you?")
    }
}