package main

import (
	"fmt"
	"strings"
)

func index() {
	splitedString := strings.Split("swiss", "")
	charMap := make(map[string]int)
	i := 0
	for _, v := range splitedString {
		charMap[v]++
		if i > 0 {
			fmt.Println(v)
		}
		i++
	}
}