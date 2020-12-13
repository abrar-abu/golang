package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {

	persons := make(map[int]person)
    persons[0] = person{"Bob", 20}
    persons[1] = person{"Alice", 30}
    persons[2] = person{"Fred", 40}
	fmt.Println("Map of Persons ",persons) 
	for key, value := range persons  {
		fmt.Println("key is ",key," Value {name:", value.name,", age:",value.age,"}")
	}

	// Confusion about this statement
	// var names map[int]string    // name map has int keys and string values
	
	// To intilise  and declare a map
	m := make(map[string]int)
	fmt.Println("Empty map",m)

	// Assign values
    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}