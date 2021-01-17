package main

import (
	"fmt"
	"errors"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func GetMinMax(flights []Flight) (int, int, error) {
	if len(flights) == 0 {
		return 0, 0, errors.New("No Flights Passed in")
	}

	var min = flights[0].Price
	var max = flights[0].Price

	for _, value := range flights {
		if value.Price < min {
			min = value.Price
		}
		if value.Price > max {
			max = value.Price
		}
	}
	return min, max, nil
}

func main() {
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
	fmt.Println(GetMinMax(flights))

	var emptyFlights []Flight

	fmt.Println(GetMinMax(emptyFlights))
	fmt.Println("Getting the Minimum and Maximum Flight Prices")
}
