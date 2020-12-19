/*

*/

package main

import (
	"fmt"
)

func main() {
	var Size, Sample int
	fmt.Scan(&Size)
	fmt.Scan(&Sample)
	array := make([][2]int, Sample)
	for i := 0; i < Sample; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&array[i][j])
		}
	}
	for i := 0; i < Sample; i++ {
	    if (array[i][0] < Size || array[i][1] < Size){
	        fmt.Println("UPLOAD ANOTHER")
	    } else {
	        if( array[i][0] == array[i][1] && (array[i][0]>=Size  && array[i][1]>=Size)){
					fmt.Println("ACCEPTED")
	        } else {
				fmt.Println("CROP IT")	            
	        }
	        
	    }
	}
}
