package elevator

type Direction int

const (
	Down Direction = iota - 1
	Stopped
	Up
)

type Car struct {
	buttons      []bool
	currentFloor int
	direction    Direction
}

func NewCar(
	numFloors int,
	currentFloor int,
	direction Direction,
	currentCalls []int,
) *Car {
	car := Car{
		buttons:      make([]bool, numFloors),
		currentFloor: currentFloor,
		direction:    direction,
	}
	for _, floor := range currentCalls {
		car.buttons[floor] = true
	}
	return &car
}

func (c *Car) Score(floor int, direction Direction) int {
	return floor
}

func (c *Car) Call(floor int) []bool {
	c.buttons[floor] = true
	return c.buttons
}
