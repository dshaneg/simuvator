package passenger

import "time"

// Shift represents the working schedule of a Passenger.
type Shift struct {
	begin    time.Time      // the time of day when the shift starts
	duration time.Duration  // the length of the shift
	days     []time.Weekday // days of the week the Passenger has begin times
}

var (
	DefaultShift = Shift{
		begin:    time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
		duration: time.Hour * 9,
		days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	EarlyShift = Shift{
		begin:    time.Date(0, 1, 1, 6, 0, 0, 0, time.Local),
		duration: time.Hour * 9,
		days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	LateShift = Shift{
		begin:    time.Date(0, 1, 1, 12, 0, 0, 0, time.Local),
		duration: time.Hour * 9,
		days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	NightShift = Shift{
		begin:    time.Date(0, 1, 1, 19, 0, 0, 0, time.Local),
		duration: time.Hour * 9,
		days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	WeekendShift = Shift{
		begin:    time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
		duration: time.Hour * 10,
		days:     []time.Weekday{time.Saturday, time.Sunday},
	}
)

func (s Shift) IsInShift(simTime time.Time) bool {
	end := s.begin.Add(s.duration)
	spansMidnight := s.begin.Day() != end.Day()

	if spansMidnight {
		if simTime.Hour() < s.begin.Hour() {
			return false
		}

		if simTime.Hour() > end.Hour() {
			return false
		}

		return true
	}
	startTime := time.Date()
	if simTime.Hour() < s.startHour {
		return false
	}

	if simTime.Hour() > p.shift.endHour {
		return false
	}

	return true
}

func (s Shift) isWorkday(simTime time.Time) bool {
	for _, day := range s.days {
		if day == simTime.Weekday() {
			return true
		}
	}
	return false
}
