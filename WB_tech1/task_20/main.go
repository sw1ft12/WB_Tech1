package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»
*/

func reverseWords(str string) string {
	words := strings.Fields(str)
	sz := len(words) - 1
	for i := 0; i < len(words)/2; i++ {
		words[i], words[sz-i] = words[sz-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Printf("%s", reverseWords(s))
}
