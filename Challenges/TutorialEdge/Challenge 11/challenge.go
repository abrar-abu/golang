package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"crypto/sha1"
	"encoding/hex"
)

// Flight struct which contains
// the origin, destination and price of a flight
type Flight struct {
  Origin      string
  Destination string
  Price       int
}

// IsSubset checks to see if the first set of
// flights is a subset of the second set of flights.
func IsSubset(first, second []Flight) bool {
	if len(second) < len(first) {
		return false
	}
	firstHashArray := getHashArray(first)
	secondHashArray := getHashArray(second)
	// fmt.Println(firstHashArray)
	// fmt.Println("**********************")
	// fmt.Println(secondHashArray)
	for _,value := range firstHashArray {
		if false == contains(secondHashArray,value){
			return false
		}
	}
	return true
}

func contains(s []string, item string) bool {
    for _, value := range s {
        if value == item {
            return true
        }
	}
    return false
}

func getHashArray(first []Flight) ([]string) {
	var hashArray = make([]string, len(first))
	for i,value := range first {
		h := sha1.New()
		h.Write(toByteArray(value))
		sha1_hash := hex.EncodeToString(h.Sum(nil))
		// fmt.Printf("%T",sha1_hash)
		// fmt.Println(value,sha1_hash)
		hashArray[i] = sha1_hash
	}
	return hashArray
}

func toByteArray(flight Flight) []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(flight)
	if err !=nil{
		fmt.Println("Unable to convert")
	}
	return res.Bytes()
}

/*
func deserialize(data []byte) *Flight {
	var flight Flight

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&flight)

	handle(err)

	return &flight
}
*/

func main() {
  fmt.Println("Sets and Subsets Challenge")
  firstFlights := []Flight{
    Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
    Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
	// Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
	// Flight{Origin: "ABU", Destination: "SNG", Price: 3000},
  }

  secondFlights := []Flight{
    Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
    Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
    Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
    Flight{Origin: "GLA", Destination: "AMS", Price: 500},
  }

  subset := IsSubset(firstFlights, secondFlights)
  fmt.Println(subset)
}
