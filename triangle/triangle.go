// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary

// Pseudo brainstorm:

// Notes:
// For a shape to be a triangle at all, all sides have to be of length > 0,
// and the sum of the lengths of any two sides must be greater than or equal to the length of the third side.

package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a == 0 || b == 0 || c == 0 {
		return "NaT"
	}
	if a+b >= c {
		if a == b && b == c && a == c {
			return "Equ"
		}
		if a == b || a == c {
			return "Iso"
		}
	}
	return "Sca"
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	// var k Kind
	// return k
}
