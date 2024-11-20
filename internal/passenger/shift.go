package passenger

import (
	"time"
)

// Shift represents the working schedule of a Passenger.
type Shift struct {
	Begin    time.Time      // the time of day when the shift starts (note that shifts crossing midnight may finish their shift on a day not in days)
	Duration time.Duration  // the length of the shift
	Days     []time.Weekday // days of the week the Passenger has begin times
}

var (
	DefaultShift = Shift{
		Begin:    time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
		Duration: time.Hour * 9,
		Days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	EarlyShift = Shift{
		Begin:    time.Date(0, 1, 1, 6, 30, 0, 0, time.Local),
		Duration: time.Hour * 9,
		Days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	LateShift = Shift{
		Begin:    time.Date(0, 1, 1, 12, 0, 0, 0, time.Local),
		Duration: time.Hour * 9,
		Days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	NightShift = Shift{
		Begin:    time.Date(0, 1, 1, 19, 0, 0, 0, time.Local),
		Duration: time.Hour * 9,
		Days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	WeekendShift = Shift{
		Begin:    time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
		Duration: time.Hour * 10,
		Days:     []time.Weekday{time.Saturday, time.Sunday},
	}
)

func (s Shift) IsInShift(simTime time.Time) bool {

	startTime, endTime := s.calculateShiftTimes(simTime)

	if !s.isWorkday(startTime) {
		return false
	}

	if simTime.Before(startTime) {
		return false
	}

	if simTime.After(endTime) {
		return false
	}

	return true
}

func (s Shift) calculateShiftTimes(simTime time.Time) (time.Time, time.Time) {
	startTime := time.Date(
		simTime.Year(),
		simTime.Month(),
		simTime.Day(),
		s.Begin.Hour(),
		s.Begin.Minute(),
		0, 0, time.Local)
	endTime := startTime.Add(s.Duration)

	if startTime.Day() != endTime.Day() {
		// back up a day on the shift, and since we cross midnight,
		// we know that if the simTime is before the end of the backed up shift
		// on the next day, we should use the shift that started yesterday
		todayEnd := endTime.Add(-24 * time.Hour)
		if simTime.Before(todayEnd) {
			return startTime.Add(-24 * time.Hour), todayEnd
		}
	}

	return startTime, endTime
}

func (s Shift) isWorkday(startTime time.Time) bool {
	weekday := startTime.Weekday()
	for _, day := range s.Days {
		if day == weekday {
			return true
		}
	}
	return false
}
