package main
import "fmt"
func main() {
    fmt.Println("Hello, world!")
    data := []int{522,5,8,9,8}
    fmt.Println(data)
    SquareAllSlice(data)
    fmt.Println(data)
    data1 := [5]int{522,5,8,9,8}
    fmt.Println(data1)
    SquareAll(&data1)
    fmt.Println(data1)
}
func SquareAllSlice (a[]int){
    for i,_ := range a {
        a[i] = a[i] *a[i]
    }
}
func SquareAll (a *[5]int){
    for i,_ := range a {
        a[i] = a[i] *a[i]
    }
}