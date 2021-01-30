package main

import (
	"encoding/json"
	"fmt"
)

type Stocks struct {
	Stocks []Stocks
}

type Stock struct {
	Ticker   string `json:"ticker"`
	Dividend float64    `json:"dividend"`
}

// HighestDividend iterates through a JSON string of stocks
// unmarshalls them into a struct and returns the Stock with
// the highest dividend
func HighestDividend(jsonString string) string {
	// implement me

	var ticker []Stock
	/*
		bytes, err := json.Marshal(json)
	    if err != nil {
	        panic(err)
		}
		fmt.Println(bytes)

	  	err := json.Unmarshal(bytes, &ticker)
	  	if err != nil {
			fmt.Println("Unable to unmarshal json") // error message
	  	}
	  	if len(ticker) == 0 {
	    	return ""
		}

		js, _ := json.Marshal(jsonString)
	*/
	test := []byte(jsonString)
	fmt.Println(test)
	fmt.Println([]byte(jsonString))
	if err := json.Unmarshal(test, &ticker); err != nil {
		panic(err)
	}

	highest := ticker[0]
	for _, v := range ticker {
		if highest.Dividend < v.Dividend {
			highest = v
		}
	}
	return highest.Ticker
}

func main() {
	fmt.Println("Stock Price AI")

	stocks_json := `[
    {"ticker":"APPL","dividend":0.5},
    {"ticker":"GOOG","dividend":0.2},
    {"ticker":"MSFT","dividend":0.3}
  ]`

	highestDividend := HighestDividend(stocks_json)
	fmt.Println(highestDividend)
}
