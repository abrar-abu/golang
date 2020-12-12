package main

import (
    "fmt"
    "time"
)

 
func t(i interface{}) string {
    switch i.(type) {
        case string:
        	return "A string value"
        case int:
			return "A number"
		case bool:
			return "A boolean"
		default:
        	return "Other"
    }
}

func f(c int) {
    switch c {
        case 2:
        fmt.Println("2")
        fallthrough
        case 4:
        fmt.Println("4")
        case 8:
        fmt.Println("8")
    }
}

func main() {
	f(2)

	var a interface{} = "A sample string"
	fmt.Println(t(a)) 		// A string value
	fmt.Println(t(12)) 		// A number value
	fmt.Println(t(true)) 	// A boolen value
	fmt.Println(t(3+5i))	// Other

    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

	t := time.Now()
	fmt.Println("time is ",t)
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}