package main

import "fmt"

func DoubleChars(original string) string {
	endString := make([]rune, 0, len(original)*2)
	for _, c := range original {
	  	endString = append(endString, c, c)
	}
	return string(endString)
}

func main() {
  fmt.Println("Smallest Difference Challenge")

  original := "gophers"
  doubled := DoubleChars(original)
  fmt.Println(doubled) // ggoopphheerrss
}
