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
		// 0 -> 1 -> 2 -> 3 (3 floors for a subscore of 3)
		// plus stopping on 2, so 5 * 1 stop
		// gives a score of 8
		name:          "On ground floor, 2 already called, call to 3 expect score 8",
		numFloors:     5,
		currentCalls:  []int{2},
		currentFloor:  0,
		carDirection:  elevator.Stopped,
		callFloor:     3,
		callDirection: elevator.Up,
		expected:      8,
	},
	{
		// 2 -> 3 -> 4 -> 3 -> 2 -> 1 -> 0 (traverse 6 floors for a subscore of 6)
		// plus stopping on 4, so 5 * 1 stop
		// gives a score of 11
		name:          "On floor 2, 4 already called, call to 0(ground) expect score 11",
		numFloors:     5,
		currentCalls:  []int{4},
		currentFloor:  2,
		carDirection:  elevator.Up,
		callFloor:     0,
		callDirection: elevator.Down, // irrelevant so far
		expected:      11,
	},
	{
		// 2 -> 1 -> 0 -> 1 -> 2 -> 3 -> 4 (traverse 6 floors for a subscore of 6)
		// plus stopping on 4, so 5 * 1 stop
		// gives a score of 11
		name:          "On floor 2, 0(ground) already called, call to 4(top) expect score 11",
		numFloors:     5,
		currentCalls:  []int{0},
		currentFloor:  2,
		carDirection:  elevator.Down,
		callFloor:     4,
		callDirection: elevator.Down,
		expected:      11,
	},
	{
		// 2 -> 2 (traverse 0 floors for a subscore of 0)
		// gives a score of 0
		name:          "On floor 2, no calls, call to 2(current floor) expect score 0",
		numFloors:     5,
		currentCalls:  []int{},
		currentFloor:  2,
		carDirection:  elevator.Stopped,
		callFloor:     2,
		callDirection: elevator.Down,
		expected:      0,
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
