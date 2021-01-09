package main

import "fmt"

// Flight - a struct that
// contains information about flights
type Flight struct {
  Origin string
  Destination string
  Price int
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
  for index,_ := range flights {
    for innerIndex := 0 ; innerIndex<len(flights)-index-1;innerIndex++ {
      if flights[innerIndex].Price > flights[innerIndex+1].Price {
        flights[innerIndex], flights[innerIndex+1] = flights[innerIndex+1],flights[innerIndex]
      }
    }
  }
  return flights
}

func printFlights(flights []Flight) {
  for _, flight := range flights {
    fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
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
