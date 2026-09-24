package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// panic("Please implement the Schedule function")
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t

}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// panic("Please implement the HasPassed function")
	// now := time.Now()
	// t, err := time.Parse(date, "July 24, 2019 13:15:00")
	// if err != nil {
	// 	panic(err)
	// }

	// return t.Before(now)

	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(t)

	return elapsed > 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// t, err := time.Parse("Thursday, July 25, 2019 13:45:00", date) (This doesn't work)
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	hour := t.Hour()

	if hour >= 12 && hour < 18 {
		return true
	}

	return false

}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// panic("Please implement the Description function")
	msg := "You have an appointment on"

	// t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	t, err := time.Parse("1/2/2006 15:04:00", date)
	outputDate := t.Format("Monday, January 2, 2006")
	outputTime := t.Format("15:04")
	if err != nil {
		panic(err)
	}

	// return msg + outputDate + " at " + outputTime + "."
	return fmt.Sprintf("%s %s, at %s.", msg, outputDate, outputTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
// => 2020-09-15 00:00:00 +0000 UTC
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
