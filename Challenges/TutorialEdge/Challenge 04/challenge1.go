package main

import "fmt"

type Developer struct {
  Name string
  Age int
}

func FilterUnique(developers []Developer) []string {
	uniqueMap := make(map[string]bool)
	var uniqueNames []string

	for _, dev := range developers {
		if _, exists := uniqueMap[dev.Name]; !exists {
			uniqueMap[dev.Name] = true
			uniqueNames = append(uniqueNames, dev.Name)
		}
	}
	return uniqueNames
}

func main() {
  fmt.Println("Filter Unique Challenge")
  developers := []Developer{
    Developer{Name: "Elliot"},
    Developer{Name: "Alan"},
    Developer{Name: "Jennifer"},
    Developer{Name: "Graham"},
    Developer{Name: "Paul"},
    Developer{Name: "Alan"},
  }
  fmt.Println(FilterUnique(developers))
}

