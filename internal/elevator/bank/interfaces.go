package bank

import "github.com/dshaneg/elevator/internal/elevator/car"

// type Scorer interface {
// 	Score(floor int, direction Direction) int
// }

// ElevatorController is an interface with two methods: Call and Move
type Member interface {
	Score(floor int, direction car.Direction) int
	Call(floor int) []bool
	Floor() int
	Direction() car.Direction
	Status() car.Status
}
