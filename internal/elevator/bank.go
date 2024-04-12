package elevator

import (
	"errors"
	"math"
)

type Bank struct {
	cars []Scorer
}

func NewBank(numFloors int, cars []Scorer) (Bank, error) {
	if len(cars) == 0 {
		return Bank{}, errors.New("elevator: NewBank requires a non-nil, non-empty slice of Scorers")
	}
	bank := Bank{
		cars: cars,
	}
	return bank, nil
}

func (b Bank) Call(floor int, direction Direction) (carIndex int) {
	lowestScore := math.MaxInt

	for i, car := range b.cars {
		score := car.Score(floor, direction)
		if score < lowestScore {
			lowestScore = score
			carIndex = i
		}
	}

	return carIndex
}
