package bank

import (
	"errors"
	"math"

	"github.com/dshaneg/elevator/internal/elevator/car"
)

// Bank represents a collection of elevator cars that are accessed from the same landing.
type Bank struct {
	cars []Member
}

// New creates a new Bank with the given number of floors and cars.
func New(numFloors int, cars []Member) (*Bank, error) {
	if len(cars) == 0 {
		return nil, errors.New("elevator: NewBank requires a non-empty slice of Scorers")
	}
	bank := Bank{
		cars: cars,
	}
	return &bank, nil
}

// LandingStatus represents the status of a landing.
//
// If any car is Loading at the landing, the status will be `Loading`.
// If a call has been made but the car has not yet arrived, the status will be `Waiting`.
// Otherwise, the status will be `Idle`.
type LandingStatus int

const (
	Idle    LandingStatus = iota // No unserved calls have been made at the landing.
	Waiting                      // A call has been made but the car has not yet arrived.
	Loading                      // At least one car is loading at the landing.
)

// Call requests an elevator car to the given floor and in the given direction.
func (b *Bank) Call(floor int, direction car.Direction) (carIndex int) {
	lowestScore := math.MaxInt

	for i, car := range b.cars {
		score := car.Score(floor, direction)
		if score < lowestScore {
			lowestScore = score
			carIndex = i
		}
	}

	b.cars[carIndex].Call(floor)

	return carIndex
}

// Status returns the status of the landing at the given floor and for the given direction.
func (b *Bank) Status(floor int, direction car.Direction) (status LandingStatus, c Member) {
	for _, c := range b.cars {
		if c.Floor() == floor && c.Direction() == direction && c.Status() == car.Loading {
			return Loading, c
		}
	}
	return Waiting, nil
}

// Car returns the car at the given index. (this feels too low-level)
func (b *Bank) Car(carIndex int) Member {
	return b.cars[carIndex]
}
