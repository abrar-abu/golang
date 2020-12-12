package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
	/*
	For refrence link 
	https://golangbyexample.com/constant-golang/
	*/

    fmt.Println(s)

	 //Unanamed untyped constant
	 fmt.Printf("Type: %T Value: %v\n", 123, 123)
	 fmt.Printf("Type: %T Value: %v\n", "circle", "circle")
	 fmt.Printf("Type: %T Value: %v\n", 5.6, 5.6)
	 fmt.Printf("Type: %T Value: %v\n", true, true)
	 fmt.Printf("Type: %T Value: %v\n", 'a', 'a')
	 fmt.Printf("Type: %T Value: %v\n", 3+5i, 3+5i)
 
	 //Named untyped constant
	 const a = 123      //Default hidden type is int
	 const b = "circle" //Default hidden type is string
	 const c = 5.6      //Default hidden type is float64
	 const d = true     //Default hidden type is bool
	 const e = 'a'      //Default hidden type is rune
	 const f = 3 + 5i   //Default hidden type is complex128
 
	 fmt.Println("")
	 fmt.Printf("Type: %T Value: %v\n", a, a)
	 fmt.Printf("Type: %T Value: %v\n", b, b)
	 fmt.Printf("Type: %T Value: %v\n", c, c)
	 fmt.Printf("Type: %T Value: %v\n", d, d)
	 fmt.Printf("Type: %T Value: %v\n", e, e)
	 fmt.Printf("Type: %T Value: %v\n", f, f)


	type myString string

	//Typed String constant
	const aa string = "abc"
	var uu = aa
	fmt.Println("Untyped named string constant")
	fmt.Printf("uu: Type: %T Value: %v\n\nn", uu, uu)

	//Below line will raise a compilation error
	//var v myString = aa

	//Untyped named string constant
	const bb = "abc"
	var ww myString = bb
	var xx = bb
	fmt.Println("Untyped named string constant")
	fmt.Printf("ww: Type: %T Value: %v\n", ww, ww)
	fmt.Printf("xx: Type: %T Value: %v\n\n", xx, xx)

	//Untyped unnamed string constant
	var yy myString = "abc"
	var zz = "abc"
	fmt.Println("Untyped unnamed string constant")
	fmt.Printf("yy: Type: %T Value: %v\n", yy, yy)
	fmt.Printf("zz: Type: %T Value: %v\n", zz, zz)


	const PI = math.Pi
	fmt.Println(PI)
    const n = 300000000

    const ans = 3e20 / n
    fmt.Println(ans)

    fmt.Println(int64(ans))

    fmt.Println(math.Sin(ans))
}