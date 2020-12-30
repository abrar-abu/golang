package main 

import (
	"encoding/csv"
	"fmt"
  	// "io"
  	"log"
  	// "strings"
  	"os"
)

type CustomerData struct {
	Customer 	string
	Area 		string
	Date 		string 
	Product 	string
	Price  		string 
	PaymentMode string
}

type ct []CustomerData

func (c *CustomerData) String() string {
	return fmt.Sprintf("{ Customer : %v \t Area : %v \t Date : %v \t Product : %v \t Price : %v \t PaymentMode : %v }\n",c.Customer,c.Area,c.Date,c.Product,c.PaymentMode)
}

func (c *ct) String() string{
	// for _,value := range c{
		fmt.Println("abu")
		fmt.Println(c)
	// }
	return "to"
}

func readFileAndGetCustomerData(fileName string) []CustomerData {
	var customers []CustomerData
	csv_file, _ := os.Open(fileName)
	csvLines, err := csv.NewReader(csv_file).ReadAll()
	if err != nil {
		log.Fatal("Issure reading csv")
	}
  	var i = 0
  	fmt.Println(i)
  	for _,line := range csvLines {
		row := CustomerData{
    		Customer: line[0],
    		Area: line[1],
    		Date: line[2],
    		Product: line[3],
    		Price: line[4],
    		PaymentMode: line[5],
		}
		fmt.Println(row.String())
		customers = append(customers, row)
    	i++
	  }
	  return customers
}

func main (){
	var fileName = "customers.csv"
	fmt.Println(readFileAndGetCustomerData(fileName))	
}