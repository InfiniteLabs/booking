package domain

import (
	"testing"
)

func TestAddBooking(t *testing.T) {
	for _, v := range addBookingTestCases {
		v.input.AddBooking(v.booking)
		if !v.input.Equals(v.expected) {
			t.Fatalf("change this message")
		}
	}
}
