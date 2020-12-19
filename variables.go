package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	Refrence link :- for data types
	https://www.tutorialspoint.com/go/go_data_types.htm
	*/

	// To intilize the string  
	// 1. Declare and initilize without data type using var
    var a = "initial"
	fmt.Println(a)
	// 2. Declare and initilize with data type using var
	var s string = "this is a string"
	fmt.Println(s)
	// 3. Declare and initilize and assign without data type without var (:=)
	t := strconv.Itoa(123)
	fmt.Println(t)
	// 4a. Declare with data type
	var testString string
	// 4b. Initilize and assign.
	testString = "test data set"
	fmt.Println(testString)

	// To initialie the numbers
    var b, c int = 1, 2
    fmt.Println(b, c)

	// To initilize the boolean
    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
	fmt.Println(f)
	
	var myChar byte = 'A'    	//Stores ASCII characters
	var myUnicode rune = '♥' 	//Stores Unicode characters
	fmt.Printf("%c = %d and %c = %U\n", myChar, myChar, myUnicode, myUnicode)

	// Default values of the data type
	// 1a. numeric
	var a1 uint8
	var a2 uint16
	var a3 uint32
	var a4 uint64
	var a5 int8
	var a6 int16
	var a7 int32
	var a8 int64
	fmt.Println(a1,a2,a3,a4,a5,a6,a7,a8)

	// 1b. numeric
	var b1 float32
	var b2 float64
	var b3 complex64
	var b4 complex128
	fmt.Println(b1,b2,b3,b4)

	// 1c. numeric
	var c1 byte
	var c2 rune
	var c3 uint
	var c4 int
	var c5 uintptr
	fmt.Println(c1,c2,c3,c4,c5)
}