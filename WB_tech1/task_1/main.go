package main

import "github.com/stretchr/testify/assert"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)
*/

type Human struct {
	a int
	b int
}

func (h Human) sum() int {
	return h.a + h.b
}

// Встраиваем структуру Human в Action. Action содержит методы и поля Human
type Action struct {
	Human
}

func NewAction() *Action {
	return &Action{
		Human{1, 2},
	}
}

func main() {
	a := NewAction()
	assert.Equal(nil, a.sum(), 3)
}
