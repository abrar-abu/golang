package main

import (
    "math/rand"
    "time"
    "fmt"
)

// Employee type
type Employee interface {
    Language() string
    Age() int
}

// Engineer type
type Engineer struct {
    Name string
}

// Language interface method for type engineer
func (e Engineer) Language() string {
    return e.Name + " programs in Go"
}


// Age function
func (e Engineer) Age() int {
    return rand.Intn(100)
}


func main() {
    // Randomize the seed used in rand.Intn to get a new int every time the program runs
    rand.Seed(time.Now().UTC().UnixNano())

    // This will throw an error
    var programmers []Employee

    // Implement interface for type Engineer
    var elliot = Engineer{Name: "Elliot"}

    programmers = append(programmers, elliot)

    fmt.Println(elliot.Name)
    fmt.Println(elliot.Age)
    fmt.Println(elliot.Language())
    fmt.Println(elliot.Age())

    
    var abu = Engineer{Name: "Abu"}

    programmers = append(programmers, abu)

    fmt.Println(abu.Name)
    fmt.Println(abu.Age)
    fmt.Println(abu.Language())
    fmt.Println(abu.Age())

    fmt.Println(programmers)
    
}