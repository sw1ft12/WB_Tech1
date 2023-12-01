package main

import (
	"math"

	"github.com/stretchr/testify/assert"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x float64
	y float64
}

func (p Point) Dist(q Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func main() {
	{
		p := NewPoint(0, 0)
		q := NewPoint(3, 4)
		assert.Equal(nil, p.Dist(*q), 5.0)
	}

	{
		p := NewPoint(1, 2)
		q := NewPoint(4, 6)
		assert.Equal(nil, p.Dist(*q), 5.0)
	}

	{
		p := NewPoint(0, 0)
		q := NewPoint(1, 1)
		assert.Equal(nil, p.Dist(*q), math.Sqrt2)
	}

	{
		p := NewPoint(5, 11)
		q := NewPoint(11, 19)
		assert.Equal(nil, p.Dist(*q), 10.0)
	}

	{
		p := NewPoint(1, 7)
		q := NewPoint(6, 12)
		assert.Equal(nil, p.Dist(*q), 5*math.Sqrt2)
	}

}
