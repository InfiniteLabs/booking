package main

import (
	"time"
)

type Room struct {
	Id          int64
	Name        string
	Description string
	Bookings    []TimeSlot `db:"-"`
}

type TimeSlot struct {
	Start time.Time
	End   time.Time
}

type TimeSlotSlice []TimeSlot

func (t TimeSlotSlice) Len() int {
	return len(t)
}

func (t TimeSlotSlice) Less(i, j int) bool {
	return t[j].Start.After(t[i].Start)
}

func (t TimeSlotSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
