package errors

import (
    "fmt"
    "time"
)


type TimeSlotOverlappingError struct {
	Start time.Time
	End time.Time
	RoomId int
}

func (e TimeSlotOverlappingError) Error() string {
	return fmt.Sprintf("Time from %s to %s in the room %s already booked", e.Start, e.End, e.RoomId)
}