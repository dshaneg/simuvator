package passenger

// Status represents the current state of a Passenger.
type Status int

// might be able to simplify this list to only Idle, Waiting and Riding.

const (
	Idle        Status = iota // not at work (outside my shift's time window).
	Active                    // at work (inside my shift's time window).
	WaitingDown               // waiting at the elevator landing to go down.
	WaitingUp                 // waiting at the elevator landing to go up.
	Riding                    // riding in an elevator car.
)
