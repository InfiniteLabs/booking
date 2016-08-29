package domain

import (
	"time"
)

var addBookingTestCases = []struct {
	input    Room
	booking  TimeSlot
	expected Room
}{
	{
		Room{
			0,
			"234",
			"some description",
			[]TimeSlot{},
		},
		TimeSlot{
			time.Date(2006, 1, 1, 12, 0, 0, 0, time.UTC),
			time.Date(2006, 1, 1, 13, 0, 0, 0, time.UTC),
		},
		Room{
			1,
			"546",
			"some  other description",
			[]TimeSlot{
				{
					time.Date(2006, 1, 1, 12, 0, 0, 0, time.UTC),
					time.Date(2006, 1, 1, 13, 0, 0, 0, time.UTC),
				},
			},
		},
	},
}
