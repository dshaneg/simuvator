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
	distance := c.findDistance(floor)
	stops := c.countStops()

	return distance + stops*5
}

func (c *Car) countStops() int {
	stops := 0
	for _, pressed := range c.buttons {
		if pressed {
			stops++
		}
	}
	return stops
}

func (c *Car) findDistance(floor int) int {
	distance := 0
	switch {
	case floor < c.currentFloor:
		if c.direction == Up {
			distance += (c.topStop() - c.currentFloor) * 2
		}
		distance += c.currentFloor - floor
	case floor > c.currentFloor:
		if c.direction == Down {
			distance += (c.currentFloor - c.bottomStop()) * 2
		}
		distance += floor - c.currentFloor
	}

	return distance
}

func (c *Car) bottomStop() int {
	for i, pressed := range c.buttons {
		if pressed {
			return i
		}
	}
	return 0
}

func (c *Car) topStop() int {
	for i := len(c.buttons) - 1; i >= 0; i-- {
		if c.buttons[i] {
			return i
		}
	}
	return len(c.buttons) - 1
}

func (c *Car) Call(floor int) []bool {
	c.buttons[floor] = true
	return c.buttons
}
