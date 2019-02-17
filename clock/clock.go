// Hint
// time module parsetime duration
package clock

import "fmt"

// PSEUDO Brainstorm
// 24 hour clock would be easier
// hours, minutes, subtractions/additions
// Needs to work with ==
// sets the base time for the clock. Returns time in 24 hour format

// !!!!! IDEA: Make a clock that revolves around minutes rather than using any functions within Go
// https://stackoverflow.com/questions/30956224/go-how-to-format-number-digit-count

// Resource: https://golangbot.com/inheritance/

// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
// (Clock) Subtract(minutes int) Clock

type Clock int

const (
	hour = 60
	day  = 24 * hour
)

// Creates a new 24 hour clock period
func New(h int, m int) Clock {
	return Clock(0).Add(hour*h + m)
}

// Takes two arguments, layout string and string to format. The layout string must use the exact times specified in the docs as a reference. Ex: parsestring(“12:15”, “00:00”)
// Returns a time object u can call .Hours() to get hours and .Minutes() to get respective hours and minutes on of the time object
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/hour, c%hour)
}

// accepts 1 arguments. Minutes to add
func (c Clock) Add(min int) Clock {
	return Clock((int(c) + day + min%day) % day)
}

// accepts 1 arguments. Minutes to subtract
func (c Clock) Subtract(min int) Clock {
	return Clock((int(c) + day - min%day) % day)
}
