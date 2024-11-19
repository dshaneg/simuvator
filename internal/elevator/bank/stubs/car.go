package stubs

import "github.com/dshaneg/elevator/internal/elevator/car"

type Car struct {
	score     int
	CallCount int
}

func NewCar(score int) *Car {
	return &Car{
		score: score,
	}
}

func (c *Car) Score(floor int, direction car.Direction) int {
	return c.score
}

func (c *Car) Call(floor int) []bool {
	c.CallCount++
	return []bool{}
}

func (c *Car) Floor() int {
	return 0
}

func (c *Car) Direction() car.Direction {
	return car.Up
}

func (c *Car) Status() car.Status {
	return car.Parked
}
