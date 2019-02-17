// Hint
// time module parsetime duration
package clock

import "time"

// PSEUDO Brainstorm
// 24 hour clock would be easier
// hours, minutes, subtractions/additions
// Needs to work with ==
// sets the base time for the clock. Returns time in 24 hour format

// Resource: https://golangbot.com/inheritance/

// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
// (Clock) Subtract(minutes int) Clock

type Clock struct {
	h int
	m int
}

// Creates a new 24 hour clock period
func New(hours int, minutes int) Clock {
	t, _ = time.Parse(time.RFC822, "2013-Feb-03")
}

// Takes two arguments, layout string and string to format. The layout string must use the exact times specified in the docs as a reference. Ex: parsestring(“12:15”, “00:00”)
// Returns a time object u can call .Hours() to get hours and .Minutes() to get respective hours and minutes on of the time object
func (c Clock) String() string {

}

// accepts 1 arguments. Minutes to add
func (c Clock) Add(minAdd int) Clock {
	c.h
}

// accepts 1 arguments. Minutes to subtract
func (c Clock) Subtract(minSub int) Clock {

}
