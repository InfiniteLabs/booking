package domain

import (
	"time"
)

type Room struct {
	Id          int
	Name        string
	Description string
	BookedSlots []Slot
}

type Slot struct {
	start time.Time
	end   time.Time
}


