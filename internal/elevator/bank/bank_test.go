package bank_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dshaneg/elevator/internal/elevator/bank"
	"github.com/dshaneg/elevator/internal/elevator/bank/stubs"
	"github.com/dshaneg/elevator/internal/elevator/car"
)

func TestCallError(t *testing.T) {
	cars := []bank.Member{}
	_, err := bank.New(5, cars)

	assert.Error(t, err)
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
			cars := []bank.Member{}
			for _, score := range tc.scores {
				cars = append(
					cars,
					stubs.NewCar(score),
				)
			}

			b, err := bank.New(tc.numFloors, cars)
			assert.NoError(t, err)

			// neither parameter matters since using a stub
			got := b.Call(0, car.Up)
			assert.Equal(t, tc.expectedCarIndex, got)
		})
	}
}

func TestCallCallsCarCall(t *testing.T) {
	c := stubs.NewCar(0)
	b, err := bank.New(5, []bank.Member{c})
	assert.NoError(t, err)

	b.Call(0, car.Up)
	assert.Equal(t, 1, c.CallCount)
}
