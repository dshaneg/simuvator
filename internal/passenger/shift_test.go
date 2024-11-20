package passenger_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dshaneg/elevator/internal/passenger"
)

var (
	OfficeShift = passenger.Shift{
		Begin:    time.Date(0, 1, 1, 8, 0, 0, 0, time.Local),
		Duration: time.Hour * 9,
		Days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	NonZeroMinutesShift = passenger.Shift{
		Begin:    time.Date(0, 1, 1, 6, 30, 0, 0, time.Local),
		Duration: time.Hour * 9,
		Days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	SpansMidnightShift = passenger.Shift{
		Begin:    time.Date(0, 1, 1, 19, 0, 0, 0, time.Local),
		Duration: time.Hour * 9,
		Days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	}

	tue0600AM = time.Date(2024, 11, 19, 06, 00, 0, 0, time.Local)
	tue1000AM = time.Date(2024, 11, 19, 10, 00, 0, 0, time.Local)
	tue0331PM = time.Date(2024, 11, 19, 15, 31, 0, 0, time.Local)
	tue0600PM = time.Date(2024, 11, 19, 18, 00, 0, 0, time.Local)
	tue0900PM = time.Date(2024, 11, 19, 21, 00, 0, 0, time.Local)
	wed0300AM = time.Date(2024, 11, 20, 03, 00, 0, 0, time.Local)
	sat0300AM = time.Date(2024, 11, 23, 03, 00, 0, 0, time.Local)
	sat1000AM = time.Date(2024, 11, 23, 10, 00, 0, 0, time.Local)
)

func TestIsInShift(t *testing.T) {
	tests := []struct {
		name     string
		shift    passenger.Shift
		simTime  time.Time
		expected bool
	}{
		{
			name:     "Returns true when the time is within the shift",
			shift:    OfficeShift,
			simTime:  tue1000AM,
			expected: true,
		},
		{
			name:     "Returns false when the time is within the shift on the WRONG day",
			shift:    OfficeShift,
			simTime:  sat1000AM,
			expected: false,
		},
		{
			name:     "Returns false when the time is before the shift on the right day",
			shift:    OfficeShift,
			simTime:  tue0600AM,
			expected: false,
		},
		{
			name:     "Returns false when the time is before the shift on the right day, within the start hour",
			shift:    NonZeroMinutesShift,
			simTime:  tue0600AM,
			expected: false,
		},
		{
			name:     "Returns false when the time is after the shift on the same day",
			shift:    OfficeShift,
			simTime:  tue0900PM,
			expected: false,
		},
		{
			name:     "Returns false when the time is after the shift on the same day, same hour",
			shift:    NonZeroMinutesShift,
			simTime:  tue0331PM,
			expected: false,
		},
		{
			name:     "Returns true when shift crosses midnight and sim time is during the shift after midnight",
			shift:    SpansMidnightShift,
			simTime:  wed0300AM,
			expected: true,
		},
		{
			name:     "Returns true when shift crosses midnight and sim time is during the shift before midnight",
			shift:    SpansMidnightShift,
			simTime:  tue0900PM,
			expected: true,
		},
		{
			name:     "Returns true when shift crosses midnight and the time is within the shift that started on the right day but finished on an off day",
			shift:    SpansMidnightShift,
			simTime:  sat0300AM,
			expected: true,
		},
		{
			name:     "Returns false when shift crosses midnight and the time is before the shift",
			shift:    SpansMidnightShift,
			simTime:  tue0600PM,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.shift.IsInShift(tc.simTime))
		})
	}
}
