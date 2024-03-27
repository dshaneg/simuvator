package elevator_test

import (
	"slices"
	"testing"

	"github.com/dshaneg/elevator/internal/elevator"
)

var scoreCases = []struct {
	name          string
	numFloors     int
	currentCalls  []int
	currentFloor  int
	carDirection  elevator.Direction
	callFloor     int
	callDirection elevator.Direction
	expected      int
}{
	{
		name:          "No calls on ground floor, call to 1 expect score 1",
		numFloors:     5,
		currentCalls:  []int{},
		currentFloor:  0,
		carDirection:  elevator.Stopped,
		callFloor:     1,
		callDirection: elevator.Up,
		expected:      1,
	},
	{
		name:          "No calls on ground floor, call to 2 expect score 2",
		numFloors:     5,
		currentCalls:  []int{},
		currentFloor:  0,
		carDirection:  elevator.Stopped,
		callFloor:     2,
		callDirection: elevator.Up,
		expected:      2,
	},
	{
		// distance between current and call floor (in floors) + 5*number of stops
		name:          "On ground floor, 2 already called, call to 3 expect score 7",
		numFloors:     5,
		currentCalls:  []int{2},
		currentFloor:  0,
		carDirection:  elevator.Stopped,
		callFloor:     3,
		callDirection: elevator.Up,
		expected:      7,
	},
}

func TestScore(t *testing.T) {
	for _, tc := range scoreCases {
		t.Run(tc.name, func(t *testing.T) {
			c := elevator.NewCar(tc.numFloors, tc.currentFloor, tc.carDirection, tc.currentCalls)

			got := c.Score(tc.callFloor, tc.callDirection)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestCallFloorSetsCallButton(t *testing.T) {
	const numFloors = 10
	c := elevator.NewCar(numFloors, 0, 0, []int{})
	calls := c.Call(2)
	expected := []bool{false, false, true, false, false, false, false, false, false, false}
	if !slices.Equal(expected, calls) {
		t.Errorf("expected %v, got %v", expected, calls)
	}
}
