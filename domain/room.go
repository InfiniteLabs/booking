package domain

import (
	"sort"
	"sync"
	"time"
	"booking/errors"
)

type Room struct {
	id          int
	name        string
	description string
	bookings    []TimeSlot
}

func (r *Room) setId(id int) {
	r.id = id
}

func (r *Room) setName(name string) {
	r.name = name
}

func (r *Room) setDescription(description string) {
	r.description = description
}

var makeSliceOnce sync.Once

func (r *Room) AddBooking(s TimeSlot) error {

	makeSliceOnce.Do(func() {
		r.bookings = make([]TimeSlot, 0, 0)
	})

	for _, v := range r.bookings {
		if v.IsOverlap(s) {
			return errors.TimeSlotOverlappingError{v.start, v.end, r.id}
		}
	}

	r.bookings = append(r.bookings, s)
	sort.Sort(TimeSlotSlice(r.bookings))

	return nil
}

func (s1 TimeSlot) IsOverlap(s2 TimeSlot) bool {
	return !s2.start.After(s1.end) || !s1.start.After(s2.end)
}

func (r Room) removeBooking() {

}

func (r1 Room) Equals(r2 Room) bool {
	if r1.id == r2.id && r1.name == r2.name && r1.description == r2.description && testEq(r1.bookings, r2.bookings) {
		return true
	}

	return false
}

func testEq(a, b []TimeSlot) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type TimeSlotSlice []TimeSlot

func (t TimeSlotSlice) Len() int {
	return len(t)
}

func (t TimeSlotSlice) Less(i, j int) bool {
	return t[j].start.After(t[i].start)
}

func (t TimeSlotSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type TimeSlot struct {
	start time.Time
	end   time.Time
}
