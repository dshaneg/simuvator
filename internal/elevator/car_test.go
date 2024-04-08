package elevator_test

import (
	"slices"
	"testing"

	"github.com/dshaneg/elevator/internal/elevator"
)

var scoreCases = []struct {
	name             string
	numFloors        int
	currentCalls     []int
	currentFloor     int
	currentDirection elevator.Direction
	currentStatus    elevator.Status
	callFloor        int
	callDirection    elevator.Direction
	expected         int
}{
	{
		// 0 -> 1
		name:             "No calls on ground floor, call to 1 expect score 1",
		numFloors:        5,
		currentCalls:     []int{},
		currentFloor:     0,
		currentDirection: elevator.Up,
		callFloor:        1,
		callDirection:    elevator.Up,
		expected:         1,
	},
	{
		// 0 -> 1 -> 2
		name:             "No calls on ground floor, call to 2 expect score 2",
		numFloors:        5,
		currentCalls:     []int{},
		currentFloor:     0,
		currentDirection: elevator.Up,
		callFloor:        2,
		callDirection:    elevator.Up,
		expected:         2,
	},
	{
		// distance between current and call floor (in floors) + 5*number of stops
		// 0 -> 1 -> 2 -> 3 (3 floors for a subscore of 3)
		// plus stopping on 2, so 5 * 1 stop
		// gives a score of 8
		name:             "On ground floor, 2 already called, call to 3 expect score 8",
		numFloors:        5,
		currentCalls:     []int{2},
		currentFloor:     0,
		currentDirection: elevator.Up,
		callFloor:        3,
		callDirection:    elevator.Up,
		expected:         8,
	},
	{
		// 2 -> 3 -> 4 -> 3 -> 2 -> 1 -> 0 (traverse 6 floors for a subscore of 6)
		// plus stopping on 4, so 5 * 1 stop
		// gives a score of 11
		name:             "On floor 2, 4 already called, call to 0(ground) expect score 11",
		numFloors:        5,
		currentCalls:     []int{4},
		currentFloor:     2,
		currentDirection: elevator.Up,
		callFloor:        0,
		callDirection:    elevator.Down, // irrelevant so far
		expected:         11,
	},
	{
		// 2 -> 1 -> 0 -> 1 -> 2 -> 3 -> 4 (traverse 6 floors for a subscore of 6)
		// plus stopping on 4, so 5 * 1 stop
		// gives a score of 11
		name:             "On floor 2, 0(ground) already called, call to 4(top) expect score 11",
		numFloors:        5,
		currentCalls:     []int{0},
		currentFloor:     2,
		currentDirection: elevator.Down,
		callFloor:        4,
		callDirection:    elevator.Down,
		expected:         11,
	},
	{
		// 2 -> 2 (traverse 0 floors for a subscore of 0)
		// gives a score of 0
		name:             "On floor 2, no calls, call to 2(current floor) expect score 0",
		numFloors:        5,
		currentCalls:     []int{},
		currentFloor:     2,
		currentDirection: elevator.Down,
		callFloor:        2,
		callDirection:    elevator.Down,
		expected:         0,
	},
}

func TestScore(t *testing.T) {
	for _, tc := range scoreCases {
		t.Run(tc.name, func(t *testing.T) {
			settings := elevator.CarSettings{
				Floor:     tc.currentFloor,
				Direction: tc.currentDirection,
				Status:    tc.currentStatus,
				Calls:     tc.currentCalls,
			}
			c := elevator.NewCar(tc.numFloors, settings)

			got := c.Score(tc.callFloor, tc.callDirection)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

var tickCases = []struct {
	name              string
	numFloors         int
	currentCalls      []int
	currentFloor      int
	currentDirection  elevator.Direction
	currentStatus     elevator.Status
	expectedFloor     int
	expectedDirection elevator.Direction
	expectedCalls     []int
	expectedStatus    elevator.Status
}{
	{
		// 2 -> 2
		name:              "No calls on second floor, expect no change",
		numFloors:         5,
		currentCalls:      []int{},
		currentFloor:      2,
		currentDirection:  elevator.Up,
		expectedFloor:     2,
		expectedDirection: elevator.Up,
		expectedCalls:     []int{},
	},
	{
		// 2 -> 3
		name:              "On floor 2, going up to floor 3, direction up",
		numFloors:         5,
		currentCalls:      []int{3},
		currentFloor:      2,
		currentDirection:  elevator.Up,
		expectedFloor:     3,
		expectedDirection: elevator.Up,
		expectedCalls:     []int{},
	},
	{
		// 2 -> 3
		name:              "On floor 2, going up to floor 3, direction opposite",
		numFloors:         5,
		currentCalls:      []int{3},
		currentFloor:      2,
		currentDirection:  elevator.Down,
		expectedFloor:     3,
		expectedDirection: elevator.Up,
		expectedCalls:     []int{},
	},
	{
		name:              "On floor 2, request floor 0, floor 4 already queued",
		numFloors:         5,
		currentCalls:      []int{0, 4},
		currentFloor:      2,
		currentDirection:  elevator.Up,
		expectedFloor:     3,
		expectedDirection: elevator.Up,
		expectedCalls:     []int{0, 4},
	},
	{
		name:              "On floor 3, request floor 0, floor 4 already queued",
		numFloors:         5,
		currentCalls:      []int{0, 4},
		currentFloor:      3,
		currentDirection:  elevator.Up,
		expectedFloor:     4,
		expectedDirection: elevator.Down,
		expectedCalls:     []int{0},
	},
	{
		name:              "On floor 1, request floor 4, floor 0 already queued",
		numFloors:         5,
		currentCalls:      []int{0, 4},
		currentFloor:      1,
		currentDirection:  elevator.Down,
		expectedFloor:     0,
		expectedDirection: elevator.Up,
		expectedCalls:     []int{4},
	},
}

func TestTick(t *testing.T) {
	for _, tc := range tickCases {
		t.Run(tc.name, func(t *testing.T) {
			settings := elevator.CarSettings{
				Floor:     tc.currentFloor,
				Direction: tc.currentDirection,
				Status:    tc.currentStatus,
				Calls:     tc.currentCalls,
			}
			c := elevator.NewCar(tc.numFloors, settings)

			c.Tick()

			if c.Floor() != tc.expectedFloor {
				t.Errorf("expected floor %v, got %v", tc.expectedFloor, c.Floor())
			}
			if c.Direction() != tc.expectedDirection {
				t.Errorf("expected direction %v, got %v", tc.expectedDirection, c.Direction())
			}
			calls := c.Calls()
			if !equalUnsorted(tc.expectedCalls, calls) {
				t.Errorf("expected calls %v, got %v", tc.expectedCalls, calls)
			}
		})
	}
}

func equalUnsorted(s1, s2 []int) bool {
	slices.Sort(s1)
	slices.Sort(s2)
	return slices.Equal(s1, s2)
}

func TestCallFloorSetsCallButton(t *testing.T) {
	const numFloors = 10
	c := elevator.NewCar(numFloors, elevator.CarSettings{})
	calls := c.Call(2)
	expected := []bool{false, false, true, false, false, false, false, false, false, false}
	if !slices.Equal(expected, calls) {
		t.Errorf("expected %v, got %v", expected, calls)
	}
}
