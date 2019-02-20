// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary

// Pseudo brainstorm:
// two pointers, x + 1, x + 2
// once one of the pointers has reached the end return the last line

// Notes:
// For want of a * the * was lost.
// And all for the want of a *.
package proverb

import "fmt"

// ["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"]

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var total_rhyme = []string{}
	if len(rhyme) >= 1 {
		for i := 0; i < len(rhyme)-1; i++ {
			total_rhyme = append(total_rhyme, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1]))
		}
		total_rhyme = append(total_rhyme, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
	}
	return total_rhyme
}
