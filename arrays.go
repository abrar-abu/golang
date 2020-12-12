package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {

    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
	fmt.Println("2d: ", twoD)
	var array [] int = rand.Perm(10)
	fmt.Println(array)
	fmt.Println(rand.Perm(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))

	rand.Seed(time.Now().Unix())
    //Pseudo Random Number Generator Functions with range

    //1. Intn(n int)
    fmt.Printf("Intn: %d\n", rand.Intn(10))

    //2. Int31n(n int32)
    fmt.Printf("Int31n: %d\n", rand.Int31n(10))

    //3. Int64n(n int32)
    fmt.Printf("Int64n: %d\n", rand.Int63n(10))

    //Pseudo Random Number Generator Functions without range.

    //1. Int()
    fmt.Printf("Int: %d\n", rand.Int())

    //2. Int31()
    fmt.Printf("Int31: %d\n", rand.Int31())

    //3. Int63()
    fmt.Printf("Int63: %d\n", rand.Int63())
 
    //4. Uint32()
    fmt.Printf("Uint32: %d\n", rand.Uint32())

    //4. Uint64()
    fmt.Printf("Uint64: %d\n", rand.Uint64())

    //Pseudo Random Number Generator Functions with range for floats

    //1. Float64()
    fmt.Printf("Float64: %f\n", rand.Float64())

    //2. Float32()
    fmt.Printf("Float32: %f\n", rand.Float32())
}