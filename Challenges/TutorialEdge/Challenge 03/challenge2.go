package main

import (
  "fmt"
  "sort"
)


// Flight - a struct that
// contains information about flights
type Flight struct {
  Origin string
  Destination string
  Price int
}

func SortByPrice(flights []Flight) []Flight {
	sort.Slice(flights, func(i, j int) bool { return flights[i].Price < flights[j].Price })
	return flights
}

func printFlights(flights []Flight) {
  for _, flight := range flights {
    fmt.Printf("Origin: %s, Destination: %s, Price: %d \n", flight.Origin, flight.Destination, flight.Price)
  }
}

func main() {
  // an empty slice of flights
  flights := []Flight{
    Flight{Price: 30},
    Flight{Price: 20},
    Flight{Price: 50},
    Flight{Price: 1000},
    Flight{Price: 130},
    Flight{Price: 20000},
    Flight{Price: 550},
    Flight{Price: 10},
  }

  sortedList := SortByPrice(flights)
  printFlights(sortedList)
}
