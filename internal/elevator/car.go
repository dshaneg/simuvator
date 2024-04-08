package elevator

type Direction int

const (
	Down Direction = iota
	Up
)

type Status int

const (
	Parked Status = iota
	Loading
	Traveling
)

type Car struct {
	buttons   []bool
	floor     int
	direction Direction
	status    Status
}

type CarSettings struct {
	Floor     int
	Direction Direction
	Status    Status
	Calls     []int
}

func NewCar(numFloors int, settings CarSettings) *Car {
	car := Car{
		buttons:   make([]bool, numFloors),
		floor:     settings.Floor,
		direction: settings.Direction,
		status:    settings.Status,
	}
	for _, floor := range settings.Calls {
		car.buttons[floor] = true
	}
	return &car
}

func (c *Car) Floor() int {
	return c.floor
}

func (c *Car) Direction() Direction {
	return c.direction
}

func (c *Car) Status() Status {
	return c.status
}

func (c *Car) Calls() []int {
	calls := []int{}
	for floor, called := range c.buttons {
		if called {
			calls = append(calls, floor)
		}
	}
	return calls
}

func (c *Car) Tick() {
	targetFloor := c.calculateTargetFloor()

	c.updateDirection(targetFloor)
	c.updateFloor(targetFloor)

	if targetFloor == c.floor {
		c.clearCall(targetFloor)

		targetFloor = c.calculateTargetFloor()
		c.updateDirection(targetFloor)
	}

}

func (c *Car) clearCall(floor int) {
	c.buttons[floor] = false
}

func (c *Car) updateFloor(targetFloor int) {
	if targetFloor > c.floor {
		c.floor++
	} else if targetFloor < c.floor {
		c.floor--
	}
}

func (c *Car) updateDirection(targetFloor int) {
	if c.direction == Up && targetFloor < c.floor {
		c.direction = Down
	} else if c.direction == Down && targetFloor > c.floor {
		c.direction = Up
	}
}

func (c *Car) calculateTargetFloor() (target int) {
	if c.direction == Up {
		target, found := c.findNextUpCall()
		if found {
			return target
		} else {
			target, found = c.findNextDownCall()
			if found {
				return target
			}
		}
	}

	// down
	target, found := c.findNextDownCall()
	if found {
		return target
	} else {
		target, found = c.findNextUpCall()
		if found {
			return target
		}
	}

	// no calls
	return c.floor
}

func (c *Car) findNextDownCall() (target int, found bool) {
	for i := c.floor; i >= 0; i-- {
		if c.buttons[i] {
			return i, true
		}
	}
	return 0, false
}

func (c *Car) findNextUpCall() (target int, found bool) {
	for i := c.floor; i < len(c.buttons); i++ {
		if c.buttons[i] {
			return i, true
		}
	}
	return 0, false
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
	case floor < c.floor:
		if c.direction == Up {
			distance += (c.topStop() - c.floor) * 2
		}
		distance += c.floor - floor
	case floor > c.floor:
		if c.direction == Down {
			distance += (c.floor - c.bottomStop()) * 2
		}
		distance += floor - c.floor
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
