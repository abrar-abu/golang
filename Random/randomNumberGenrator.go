package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func main() {   
     
    rand.Seed(time.Now().UnixNano())
    v := rand.Perm(8)
    fmt.Println(v)  
	
	var array [5]int
    rand.Seed(time.Now().UnixNano())
    for i:=0; i<5; i++ {
        array[i] = rand.Intn(100)
    }
    fmt.Println(array) 
}