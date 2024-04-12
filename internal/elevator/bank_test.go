package elevator_test

import (
	"testing"

	"github.com/dshaneg/elevator/internal/elevator"
)

type CarStub struct {
	score int
}

func (car CarStub) Score(floor int, direction elevator.Direction) int {
	return car.score
}

func TestCallError(t *testing.T) {
	cars := []elevator.Scorer{}
	_, err := elevator.NewBank(5, cars)

	if err == nil {
		t.Error("expected an error but did not get one")
	}
}

var bankCases = []struct {
	name             string
	numFloors        int
	scores           []int
	expectedCarIndex int
}{
	{
		name:             "Returns 0th index when it has the lowest score",
		numFloors:        5,
		scores:           []int{0, 10, 20, 100},
		expectedCarIndex: 0,
	},
	{
		name:             "Returns first lowest scored index when not 0th index",
		numFloors:        5,
		scores:           []int{100, 80, 10, 90},
		expectedCarIndex: 2,
	},
	{
		name:             "Works with negative scores",
		numFloors:        5,
		scores:           []int{100, 80, 10, -90},
		expectedCarIndex: 3,
	},
	{
		name:             "Returns first lowest scored index when there's a tie",
		numFloors:        5,
		scores:           []int{100, 10, 10, 90},
		expectedCarIndex: 1,
	},
}

func TestCall(t *testing.T) {
	for _, tc := range bankCases {
		t.Run(tc.name, func(t *testing.T) {
			cars := []elevator.Scorer{}
			for _, score := range tc.scores {
				cars = append(cars, CarStub{score})
			}

			b, err := elevator.NewBank(tc.numFloors, cars)
			if err != nil {
				t.Errorf("did not expect an error, but got %v", err)
			}

			// neither parameter matters since using a stub
			got := b.Call(0, elevator.Up)
			if got != tc.expectedCarIndex {
				t.Errorf("expected %v, got %v", tc.expectedCarIndex, got)
			}
		})
	}
}
