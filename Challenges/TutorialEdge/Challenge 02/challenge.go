package main

import (
  "fmt"
)

// Developer struct
type Developer struct {
  Name string
  Age int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
  developer :=  Developer{
	Name: name.(string),
	Age: age.(int),
  }
  return developer
}

func main() {
  fmt.Println("Hello World")

  var name interface{} = "Elliot"
  var age interface{} = 26

  dynamicDev := GetDeveloper(name, age)
  fmt.Println(dynamicDev.Name)
}