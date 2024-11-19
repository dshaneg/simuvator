package passenger

import (
	"time"

	"github.com/dshaneg/elevator/internal/elevator/bank"
	"github.com/dshaneg/elevator/internal/elevator/car"
)

// Passenger represents a person who rides our elevators.
type Passenger struct {
	bank         *bank.Bank
	primaryFloor int
	shift        Shift
	floor        int
	status       Status
	destFloor    int
	car          bank.Member
}

// Option is a functional option type that allows us to configure the Passenger.
type Option func(*Passenger)

// New creates a new Passenger with the given primary floor and options.
func New(b *bank.Bank, options ...Option) *Passenger {
	p := Passenger{
		bank:   b,
		shift:  DefaultShift,
		status: Idle,
	}

	for _, opt := range options {
		opt(&p)
	}

	return &p
}

func WithPrimaryFloor(floor int) Option {
	return func(p *Passenger) {
		p.primaryFloor = floor
	}
}

func WithShift(s Shift) Option {
	return func(p *Passenger) {
		p.shift = s
	}
}

func WithFloor(floor int) Option {
	return func(p *Passenger) {
		p.floor = floor
	}
}

func WithStatus(status Status) Option {
	return func(p *Passenger) {
		p.status = status
	}
}

func (p *Passenger) Floor() int {
	return p.floor
}

func (p *Passenger) Status() Status {
	return p.status
}

func (p *Passenger) Tick(simTime time.Time) {
	// consider implementing a state machine to manage these transitions
	// at the transition to Idle or Active, we should determine the time
	// and destination for the next elevator ride and queue it up

	isInShift := p.shift.IsInShift(simTime)
	switch {
	// coming to work
	// Idle -> WaitingUp or WaitingDown
	case p.status == Idle && isInShift && p.floor != p.primaryFloor:
		p.destFloor = p.primaryFloor
		p.call(p.destFloor)
	// done for the day
	// Active -> WaitingUp or WaitingDown
	case p.status == Active && !isInShift && p.floor != 0:
		p.destFloor = 0
		p.call(p.destFloor)
	// going up or going down
	// WaitingUp or WaitingDown -> Riding
	case p.status == WaitingUp || p.status == WaitingDown:
		p.ride()
	// exiting elevator
	// Riding -> Idle or Active
	case p.status == Riding && p.car.Floor() == p.destFloor && p.car.Status() == car.Loading:
		p.car = nil
		if isInShift {
			p.status = Active
		} else {
			p.status = Idle
		}
	}
}

func (p *Passenger) ride() {
	direction := car.Down
	if p.status == WaitingUp {
		direction = car.Up
	}
	// one of the cars in the bank may be loading, but headed the wrong direction
	// so we need to check the status in the direction we want to go
	status, c := p.bank.Status(p.floor, direction)
	if status == bank.Loading {
		p.status = Riding
		p.car = c
		p.car.Call(p.destFloor)
	}
}

func (p *Passenger) call(dest int) {
	if p.floor < dest {
		p.bank.Call(p.floor, car.Up)
		p.status = WaitingUp
		return
	}

	p.bank.Call(p.floor, car.Down)
	p.status = WaitingDown
}
