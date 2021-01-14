package main

import "fmt"

type Developer struct {
  Name string
  Age int
}

func FilterUnique(developers []Developer) []string {
  var uniques []string
  check:=make(map[string]int)
  for _,developer:=range developers{
  	check[developer.Name]=1
  } 
  for name,_ :=range check{
    uniques=append(uniques,name)
  }
  return uniques
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

