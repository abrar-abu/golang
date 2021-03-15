package main

import "fmt"

func addBorder(picture []string) []string {
	var maxLength = len(picture[0])
	for _, value := range picture {
		if len(value) > maxLength {
			maxLength = len(value)
		}
	}
	maxLength = maxLength + 2
	// fmt.Println(maxLength)
	pictureFrame := make([]string, 0)
	pictureFrame = append(pictureFrame, createString(maxLength))
	for _, value := range picture {
		pictureFrame = append(pictureFrame, "*"+value+"*")
	}
	pictureFrame = append(pictureFrame, createString(maxLength))
	return pictureFrame
}

func createString(maxLength int) string {
	var str string
	for index := 0; index < maxLength; index++ {
		str = "*" + str
	}
	return str
}
func main() {
	fmt.Printf("%v", addBorder([]string{"abc",
		"ded"}))
}
