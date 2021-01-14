package main

import "fmt"

type Developer struct {
  Name string
  Age int
}

func FilterUnique(developers []Developer) []string {
  var uniqueDevelopers []string
  for _, value := range developers {
		isPresent := false
		for _, name := range uniqueDevelopers {
			if name == value.Name {
				isPresent = true
			}
		}
		if !isPresent {
			uniqueDevelopers = append(uniqueDevelopers, value.Name)
		}
	}
  return uniqueDevelopers
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
