package main

import (
	"fmt"
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	var result string

	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == ']' || s[i] == '{' || s[i] == '}' {
			result+=string(s[i])
		} else if s[i] >= 97 && s[i] <= 123 {
			index := (s[i] - 97) + byte(k)
			result+= string(byte( (index%26)+byte(97) ))
		} else if s[i] >= 65 && s[i] <= 91 {
			index := (s[i] - 65) + byte(k)
			result+= string(byte( (index%26)+byte(65) ))
		} else {
			result+=string(s[i])
		}
	} //for

	return result
}

func main() {
	var str string
	var k int32
	fmt.Println("Enter a string : ")
	fmt.Scan(&str)
	fmt.Println("Enter key : ")
	fmt.Scan(&k)
	result := caesarCipher(str, k)
	fmt.Printf("%s\n", result)
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var ret string 
// 	// var b byte = 'a'
// 	for i := 0; i < 26; i++ {
// 		ret += string(byte(i + 97))
// 	}
// 	fmt.Println(ret)
// }
