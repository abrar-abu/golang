package main

import (
	"encoding/base64"
	"fmt"
)

func DecodeSecret(message string) string {
	data, err := base64.StdEncoding.DecodeString(message)
	fmt.Println(data)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	decoded := make([]rune, 0, len(data))

	for _, c := range data {
		decoded = append(decoded, rune(int(c)-1))
	}
	fmt.Println(decoded)
	return string(decoded)
}

func main() {
	fmt.Println("Decode the Secret")

	message := "VEZEU0ZVVFVTSk9I"
	result := DecodeSecret(message)
	fmt.Println(result)

}
