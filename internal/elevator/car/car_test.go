package car_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dshaneg/elevator/internal/elevator/car"
)

var scoreCases = []struct {
	name             string
	numFloors        int
	currentCalls     []int
	currentFloor     int
	currentDirection car.Direction
	currentStatus    car.Status
	callFloor        int
	callDirection    car.Direction
	expected         int
}{
	{
		// 0 -> 1
		name:             "No calls on ground floor, call to 1 expect score 1",
		numFloors:        5,
		currentCalls:     []int{},
		currentFloor:     0,
		currentDirection: car.Up,
		callFloor:        1,
		callDirection:    car.Up,
		expected:         1,
	},
	{
		// 0 -> 1 -> 2
		name:             "No calls on ground floor, call to 2 expect score 2",
		numFloors:        5,
		currentCalls:     []int{},
		currentFloor:     0,
		currentDirection: car.Up,
		callFloor:        2,
		callDirection:    car.Up,
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
		currentDirection: car.Up,
		callFloor:        3,
		callDirection:    car.Up,
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
		currentDirection: car.Up,
		callFloor:        0,
		callDirection:    car.Down, // irrelevant so far
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
		currentDirection: car.Down,
		callFloor:        4,
		callDirection:    car.Down,
		expected:         11,
	},
	{
		// 2 -> 2 (traverse 0 floors for a subscore of 0)
		// gives a score of 0
		name:             "On floor 2, no calls, call to 2(current floor) expect score 0",
		numFloors:        5,
		currentCalls:     []int{},
		currentFloor:     2,
		currentDirection: car.Down,
		callFloor:        2,
		callDirection:    car.Down,
		expected:         0,
	},
}

func TestScore(t *testing.T) {
	for _, tc := range scoreCases {
		t.Run(tc.name, func(t *testing.T) {
			c := car.NewCar(tc.numFloors,
				car.WithFloor(tc.currentFloor),
				car.WithDirection(tc.currentDirection),
				car.WithStatus(tc.currentStatus),
				car.WithCalls(tc.currentCalls),
			)

			got := c.Score(tc.callFloor, tc.callDirection)
			assert.Equal(t, tc.expected, got)
		})
	}
}

var tickCases = []struct {
	name              string
	numFloors         int
	currentCalls      []int
	currentFloor      int
	currentDirection  car.Direction
	currentStatus     car.Status
	expectedFloor     int
	expectedDirection car.Direction
	expectedCalls     []int
	expectedStatus    car.Status
}{
	{
		// 2 -> 2
		name:              "No calls on second floor, expect no change",
		numFloors:         5,
		currentCalls:      []int{},
		currentFloor:      2,
		currentDirection:  car.Up,
		expectedFloor:     2,
		expectedDirection: car.Up,
		expectedCalls:     []int{},
	},
	{
		// 2 -> 3
		name:              "On floor 2, going up to floor 3, direction up",
		numFloors:         5,
		currentCalls:      []int{3},
		currentFloor:      2,
		currentDirection:  car.Up,
		expectedFloor:     3,
		expectedDirection: car.Up,
		expectedCalls:     []int{},
	},
	{
		// 2 -> 3
		name:              "On floor 2, going up to floor 3, direction opposite",
		numFloors:         5,
		currentCalls:      []int{3},
		currentFloor:      2,
		currentDirection:  car.Down,
		expectedFloor:     3,
		expectedDirection: car.Up,
		expectedCalls:     []int{},
	},
	{
		name:              "On floor 2, request floor 0, floor 4 already queued",
		numFloors:         5,
		currentCalls:      []int{0, 4},
		currentFloor:      2,
		currentDirection:  car.Up,
		expectedFloor:     3,
		expectedDirection: car.Up,
		expectedCalls:     []int{0, 4},
	},
	{
		name:              "On floor 3, request floor 0, floor 4 already queued",
		numFloors:         5,
		currentCalls:      []int{0, 4},
		currentFloor:      3,
		currentDirection:  car.Up,
		expectedFloor:     4,
		expectedDirection: car.Down,
		expectedCalls:     []int{0},
	},
	{
		name:              "On floor 1, request floor 4, floor 0 already queued",
		numFloors:         5,
		currentCalls:      []int{0, 4},
		currentFloor:      1,
		currentDirection:  car.Down,
		expectedFloor:     0,
		expectedDirection: car.Up,
		expectedCalls:     []int{4},
	},
}

func TestTick(t *testing.T) {
	for _, tc := range tickCases {
		t.Run(tc.name, func(t *testing.T) {
			c := car.NewCar(tc.numFloors,
				car.WithFloor(tc.currentFloor),
				car.WithDirection(tc.currentDirection),
				car.WithStatus(tc.currentStatus),
				car.WithCalls(tc.currentCalls),
			)

			c.Tick()

			assert.Equal(t, tc.expectedFloor, c.Floor())
			assert.Equal(t, tc.expectedDirection, c.Direction())
			calls := c.Calls()
			assert.True(t, equalUnsorted(tc.expectedCalls, calls))
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
	c := car.NewCar(numFloors)
	calls := c.Call(2)
	expected := []bool{false, false, true, false, false, false, false, false, false, false}
	assert.Equal(t, expected, calls)
}
