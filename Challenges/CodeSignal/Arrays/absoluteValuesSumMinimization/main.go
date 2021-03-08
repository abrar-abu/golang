package main

func absoluteValuesSumMinimization(a []int) int {
	mapSummation := make(map[int]int)
	for index, _ := range a {
		var sum int = 0
		for _, value := range a {
			sum = sum + Abs(a[index]-value)
		}
		mapSummation[a[index]] = sum

	}
	minSum := mapSummation[a[0]]
	var key int = a[0]
	for key1, value := range mapSummation {
		if value < minSum {
			minSum = value
			key = key1
		}
	}
	return key

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Printf("%v", absoluteValuesSumMinimization([]int{2, 4, 7}))
}
