package car

// Direction is an enum type that represents the current direction of the [Car].
type Direction int

const (
	Down Direction = -1
	Up   Direction = 1
)

// Status is an enum type that represents the current status of the Car.
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

// Option is a functional option type that allows us to configure the Car
type Option func(*Car)

func NewCar(numFloors int, options ...Option) *Car {
	car := Car{
		buttons:   make([]bool, numFloors),
		floor:     0,
		direction: Up,
		status:    Parked,
	}

	for _, opt := range options {
		opt(&car)
	}

	return &car
}

// WithFloor is a functional option that sets the current floor of the Car.
func WithFloor(currentFloor int) Option {
	return func(c *Car) {
		c.floor = currentFloor
	}
}

// WithDirection is a functional option that sets the current direction of the Car.
func WithDirection(direction Direction) Option {
	return func(c *Car) {
		c.direction = direction
	}
}

// WithStatus is a functional option that sets the current status of the Car.
func WithStatus(status Status) Option {
	return func(c *Car) {
		c.status = status
	}
}

// WithCalls is a functional option that sets the current calls of the Car.
func WithCalls(calls []int) Option {
	return func(c *Car) {
		for _, floor := range calls {
			c.buttons[floor] = true
		}
	}
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

func (c *Car) Call(floor int) []bool {
	c.buttons[floor] = true
	return c.buttons
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
