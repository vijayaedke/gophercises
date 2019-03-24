package main

import (
	"fmt"
	"strings"
)

// Complete the camelcase function below.
func camelcase(s string) int32 {
	count := int32(1)
	length := len(s)
	if length <= 0 {
		return 0
	}

	for i := 0; i < length; i++ {
		if s[i] >= 65 && s[i] <= 91 {
			count += int32(1)
		}
	}

	return count

} //camelcase

func main() {
	var str string
	fmt.Println("Enter a string : ")
	fmt.Scan(&str)
	result := camelcase(strings.Trim(str, "\n\r"))
	fmt.Printf("%d\n", result)
}
