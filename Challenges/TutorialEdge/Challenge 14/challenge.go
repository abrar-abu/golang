package main

import "fmt"

func CheckLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
  fmt.Println("Check Leap Year")

  year := 2020
  fmt.Println(CheckLeapYear(year))
  fmt.Println(CheckLeapYear(1900))
  fmt.Println(CheckLeapYear(2000))
}
