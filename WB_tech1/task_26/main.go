package main

import (
	"strings"

	"github.com/stretchr/testify/assert"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой
*/

func isUnique(s string) bool {
	m := make(map[rune]bool)
	for _, c := range strings.ToLower(s) {
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}

func main() {
	assert.Equal(nil, isUnique("abcd"), true)
	assert.Equal(nil, isUnique("abCdefAaf"), false)
	assert.Equal(nil, isUnique("aabcd"), false)
	assert.Equal(nil, isUnique("aA"), false)
	assert.Equal(nil, isUnique("aBcDeF"), true)
	assert.Equal(nil, isUnique("12345"), true)
}
