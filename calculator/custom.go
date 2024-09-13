package calculator

import (
	"fmt"
	// "strings"
)

func ShouldBeEven(actual interface{}, expected ...interface{}) string {
	if len(expected) != 0 {
		return "ShouldBeEven assertion does not accept expected arguments"
	}

	number, ok := actual.(int)
	if !ok {
		return "ShouldBeEven assertion expects an integer"
	}

	if number%2 != 0 {
		return fmt.Sprintf("Expected %d to be even", number)
	}

	return "" // Assertion passed
}

func ShouldReturnFirstRepeatletter(actual string, expected string) string {
	if expected == "" {
		return "Expect a character to return as a result"
	}

	if actual == expected {
		return ""
	}

	return "Not a valid character"

}

// func ReturnFirstNonRepeatletter(str string) string {
// 	splitedString := strings.Split(str, "")
// 	charMap := make(map[string]int)
// 	i:= 0
// 	for _, v := range splitedString {
// 		charMap[v]++
// 		if i > 0 {
// 			return v
// 		}
// 		i++
// 	}
// 	return ""
// }